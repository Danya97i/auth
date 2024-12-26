package app

import (
	"context"
	"strings"

	"github.com/Danya97i/platform_common/pkg/closer"
	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Danya97i/platform_common/pkg/db/pg"
	"github.com/Danya97i/platform_common/pkg/db/transaction"
	"github.com/IBM/sarama"
	redigo "github.com/gomodule/redigo/redis"

	accessServer "github.com/Danya97i/auth/internal/api/access"
	authServer "github.com/Danya97i/auth/internal/api/auth"
	userServer "github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/client/cache"
	"github.com/Danya97i/auth/internal/client/cache/redis"
	"github.com/Danya97i/auth/internal/client/kafka"
	"github.com/Danya97i/auth/internal/client/kafka/producer"
	"github.com/Danya97i/auth/internal/config"
	"github.com/Danya97i/auth/internal/config/env"
	"github.com/Danya97i/auth/internal/repository"
	accessRuleRepo "github.com/Danya97i/auth/internal/repository/access_rule"
	logRepo "github.com/Danya97i/auth/internal/repository/logs"
	pgUserRepo "github.com/Danya97i/auth/internal/repository/user/pg"
	redisUserRepo "github.com/Danya97i/auth/internal/repository/user/redis"
	"github.com/Danya97i/auth/internal/service"
	accessService "github.com/Danya97i/auth/internal/service/access"
	authService "github.com/Danya97i/auth/internal/service/auth"
	userService "github.com/Danya97i/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig           config.PGConfig
	grpcConfig         config.GRPCConfig
	redisConfig        config.RedisConfig
	gatewayConfig      config.GatewayConfig
	swaggerConfig      config.SwaggerConfig
	kafkaConfig        config.KafkaConfig
	accessTokenConfig  config.TokenConfig
	refreshTokenConfig config.TokenConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	syncProducer sarama.SyncProducer

	kafkaProducer kafka.Producer

	userCache repository.UserCache

	userRepository repository.UserRepository

	logRepository repository.LogRepository

	accessRuleRepository repository.AccessRuleRepository

	userService service.UserService

	userServer *userServer.Server

	authServer *authServer.Server

	accessServer *accessServer.Server

	authService service.AuthService

	accessService service.AccessService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig returns pg config
func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		config, err := env.NewPgConfig()
		if err != nil {
			panic(err)
		}
		sp.pgConfig = config
	}
	return sp.pgConfig
}

// GRPCConfig returns grpc config
func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		config, err := env.NewGrpcConfig()
		if err != nil {
			panic(err)
		}
		sp.grpcConfig = config
	}
	return sp.grpcConfig
}

// RedisConfig returns redis config
func (sp *serviceProvider) RedisConfig() config.RedisConfig {
	if sp.redisConfig == nil {
		config, err := env.NewRedisConfig()
		if err != nil {
			panic(err)
		}
		sp.redisConfig = config
	}
	return sp.redisConfig
}

// GatewayConfig returns gateway config
func (sp *serviceProvider) GatewayConfig() config.GatewayConfig {
	if sp.gatewayConfig == nil {
		config, err := env.NewGatewayConfig()
		if err != nil {
			panic(err)
		}
		sp.gatewayConfig = config
	}
	return sp.gatewayConfig
}

// SwaggerConfig returns swagger config
func (sp *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if sp.swaggerConfig == nil {
		config, err := env.NewSwaggerConfig()
		if err != nil {
			panic(err)
		}
		sp.swaggerConfig = config
	}
	return sp.swaggerConfig
}

// KafkaConfig returns kafka config
func (sp *serviceProvider) KafkaConfig() config.KafkaConfig {
	if sp.kafkaConfig == nil {
		config, err := env.NewKafkaConfig()
		if err != nil {
			panic(err)
		}
		sp.kafkaConfig = config
	}
	return sp.kafkaConfig
}

// AccessTokenConfig returns access token config
func (sp *serviceProvider) AccessTokenConfig() config.TokenConfig {
	if sp.accessTokenConfig == nil {
		config, err := env.NewAccessTokenConfig()
		if err != nil {
			panic(err)
		}
		sp.accessTokenConfig = config
	}
	return sp.accessTokenConfig

}

// RefreshTokenConfig returns refresh token config
func (sp *serviceProvider) RefreshTokenConfig() config.TokenConfig {
	if sp.refreshTokenConfig == nil {
		config, err := env.NewRefreshTokenConfig()
		if err != nil {
			panic(err)
		}
		sp.refreshTokenConfig = config
	}
	return sp.refreshTokenConfig
}

// DBClient returns db client
func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		client, err := pg.NewPGClient(ctx, sp.PGConfig().DSN())
		if err != nil {
			panic(err)
		}
		if err := client.DB().Ping(ctx); err != nil {
			panic(err)
		}
		closer.Add(client.Close)
		sp.dbClient = client
	}
	return sp.dbClient
}

// TxManager returns tx manager
func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}
	return sp.txManager
}

// RedisPool returns redis pool
func (sp *serviceProvider) RedisPool() *redigo.Pool {
	if sp.redisPool == nil {
		sp.redisPool = &redigo.Pool{
			MaxIdle:     sp.RedisConfig().MaxIdle(),
			IdleTimeout: sp.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", sp.RedisConfig().Address())
			},
		}
	}

	return sp.redisPool
}

// RedisClient returns redis client
func (sp *serviceProvider) RedisClient() cache.RedisClient {
	if sp.redisClient == nil {
		sp.redisClient = redis.NewClient(sp.RedisPool(), sp.RedisConfig())
	}

	return sp.redisClient
}

// SyncProducer returns sync producer
func (sp *serviceProvider) SyncProducer() sarama.SyncProducer {
	if sp.syncProducer == nil {
		producerConfig := sarama.NewConfig()
		producerConfig.Producer.RequiredAcks = sarama.WaitForAll
		producerConfig.Producer.Retry.Max = sp.KafkaConfig().MaxRetryCount()
		producerConfig.Producer.Return.Successes = true

		syncProducer, err := sarama.NewSyncProducer(strings.Split(sp.KafkaConfig().Hosts(), ","), producerConfig)
		if err != nil {
			panic(err)
		}
		sp.syncProducer = syncProducer
	}
	return sp.syncProducer

}

// UserProducer returns user producer
func (sp *serviceProvider) UserProducer(_ context.Context) kafka.Producer {
	if sp.kafkaProducer == nil {
		sp.kafkaProducer = producer.NewProducer(sp.SyncProducer(), sp.KafkaConfig().UserTopic())
		closer.Add(sp.kafkaProducer.Close)
	}
	return sp.kafkaProducer
}

// UserRepository returns user repository
func (sp *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = pgUserRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.userRepository
}

// UserCache returns user cache
func (sp *serviceProvider) UserCache(_ context.Context) repository.UserCache {
	if sp.userCache == nil {
		sp.userCache = redisUserRepo.NewRepositoty(sp.RedisClient())
	}
	return sp.userCache
}

// UserService returns user service
func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(
			sp.UserRepository(ctx),
			sp.LogRepository(ctx),
			sp.TxManager(ctx),
			sp.UserCache(ctx),
			sp.UserProducer(ctx),
		)
	}
	return sp.userService
}

// UserServer returns user server
func (sp *serviceProvider) UserServer(ctx context.Context) *userServer.Server {
	if sp.userServer == nil {
		sp.userServer = userServer.NewServer(sp.UserService(ctx))
	}
	return sp.userServer
}

// LogRepository returns log repository
func (sp *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if sp.logRepository == nil {
		sp.logRepository = logRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.logRepository
}

// AccessRuleRepository returns access rule repository
func (sp *serviceProvider) AccessRuleRepository(ctx context.Context) repository.AccessRuleRepository {
	if sp.accessRuleRepository == nil {
		sp.accessRuleRepository = accessRuleRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.accessRuleRepository
}

//

// AuthServer returns auth server
func (sp *serviceProvider) AuthServer(ctx context.Context) *authServer.Server {
	if sp.authServer == nil {
		sp.authServer = authServer.NewServer(sp.AuthService(ctx))
	}
	return sp.authServer
}

// AccessServer returns access server
func (sp *serviceProvider) AccessServer(ctx context.Context) *accessServer.Server {
	if sp.accessServer == nil {
		sp.accessServer = accessServer.NewServer(sp.AccessService(ctx))
	}
	return sp.accessServer
}

// AuthService returns auth service
func (sp *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if sp.authService == nil {
		sp.authService = authService.NewService(
			sp.UserRepository(ctx),
			sp.AccessTokenConfig().Secret(),
			sp.RefreshTokenConfig().Secret(),
			sp.AccessTokenConfig().Expiration(),
			sp.RefreshTokenConfig().Expiration(),
		)
	}
	return sp.authService
}

// AccessService returns access service
func (sp *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if sp.accessService == nil {
		sp.accessService = accessService.NewService(
			sp.AccessRuleRepository(ctx),
			sp.AccessTokenConfig().Secret(),
		)
	}
	return sp.accessService
}

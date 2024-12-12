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

	userServer "github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/client/cache"
	"github.com/Danya97i/auth/internal/client/cache/redis"
	"github.com/Danya97i/auth/internal/client/kafka"
	"github.com/Danya97i/auth/internal/client/kafka/producer"
	"github.com/Danya97i/auth/internal/config"
	"github.com/Danya97i/auth/internal/config/env"
	"github.com/Danya97i/auth/internal/repository"
	logRepo "github.com/Danya97i/auth/internal/repository/logs"
	pgUserRepo "github.com/Danya97i/auth/internal/repository/user/pg"
	redisUserRepo "github.com/Danya97i/auth/internal/repository/user/redis"
	"github.com/Danya97i/auth/internal/service"
	userService "github.com/Danya97i/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	redisConfig   config.RedisConfig
	gatewayConfig config.GatewayConfig
	swaggerConfig config.SwaggerConfig
	kafkaConfig   config.KafkaConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	syncProducer sarama.SyncProducer

	kafkaProducer kafka.Producer

	userCache repository.UserCache

	userRepository repository.UserRepository

	userService service.UserService

	userServer *userServer.Server

	logRepository repository.LogRepository
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

func (sp *serviceProvider) RedisClient() cache.RedisClient {
	if sp.redisClient == nil {
		sp.redisClient = redis.NewClient(sp.RedisPool(), sp.RedisConfig())
	}

	return sp.redisClient
}

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

func (sp *serviceProvider) UserProducer(ctx context.Context) kafka.Producer {
	if sp.kafkaProducer == nil {
		sp.kafkaProducer = producer.NewProducer(sp.SyncProducer(), sp.KafkaConfig().UserTopic())
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

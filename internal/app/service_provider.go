package app

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/closer"
	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Danya97i/platform_common/pkg/db/pg"
	"github.com/Danya97i/platform_common/pkg/db/transaction"
	redigo "github.com/gomodule/redigo/redis"

	userServer "github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/client/cache"
	"github.com/Danya97i/auth/internal/client/cache/redis"
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
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

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
			sp.UserRepository(ctx), sp.LogRepository(ctx), sp.TxManager(ctx), sp.UserCache(ctx),
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

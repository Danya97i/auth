package app

import (
	"context"

	userServer "github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/client/db"
	"github.com/Danya97i/auth/internal/client/db/pg"
	"github.com/Danya97i/auth/internal/client/db/transaction"
	"github.com/Danya97i/auth/internal/closer"
	"github.com/Danya97i/auth/internal/config"
	"github.com/Danya97i/auth/internal/config/env"
	"github.com/Danya97i/auth/internal/repository"
	logRepo "github.com/Danya97i/auth/internal/repository/logs"
	userRepo "github.com/Danya97i/auth/internal/repository/user"
	"github.com/Danya97i/auth/internal/service"
	userService "github.com/Danya97i/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig       config.PGConfig
	grpcConfig     config.GRPCConfig
	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository
	userService    service.UserService
	userServer     *userServer.Server
	logRepository  repository.LogRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

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

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}
	return sp.txManager
}

func (sp *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.userRepository
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(
			sp.UserRepository(ctx), sp.LogRepository(ctx), sp.TxManager(ctx),
		)
	}
	return sp.userService
}

func (sp *serviceProvider) UserServer(ctx context.Context) *userServer.Server {
	if sp.userServer == nil {
		sp.userServer = userServer.NewServer(sp.UserService(ctx))
	}
	return sp.userServer
}

func (sp *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if sp.logRepository == nil {
		sp.logRepository = logRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.logRepository
}

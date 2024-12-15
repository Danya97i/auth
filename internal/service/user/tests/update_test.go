package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Danya97i/platform_common/pkg/db"
	dbMocks "github.com/Danya97i/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Danya97i/auth/internal/client/kafka"
	kafkaMocks "github.com/Danya97i/auth/internal/client/kafka/mocks"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/repository"
	repoMocks "github.com/Danya97i/auth/internal/repository/mocks"
	"github.com/Danya97i/auth/internal/service/user"
)

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoyMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txRepositoryMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache
	type userProducerMockFunc func(mc *minimock.Controller) kafka.Producer

	type args struct {
		ctx      context.Context
		id       int64
		userInfo *models.UserInfo
	}

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		userInfo = models.UserInfo{
			Name:  &name,
			Email: email,
			Role:  consts.User,
		}

		logInfo = models.LogInfo{UserID: id, Action: consts.ActionUpdate}

		nilInfoErr  = errors.New("userInfo is empty")
		userRepoErr = errors.New("user repository error")
		logRepoErr  = errors.New("log repository error")
	)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoyMock   logRepositoyMockFunc
		txRepositoryMock   txRepositoryMockFunc
		userCacheMock      userCacheMockFunc
		userProducerMock   userProducerMockFunc
	}{{
		name: "user service: update user: success case",
		args: args{
			ctx:      ctx,
			id:       id,
			userInfo: &userInfo,
		},
		err: nil,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.UpdateMock.Expect(ctx, id, userInfo).Return(nil)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			mock.SaveMock.Expect(ctx, logInfo).Return(nil)
			return mock
		},

		txRepositoryMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},

		userCacheMock: func(mc *minimock.Controller) repository.UserCache {
			mock := repoMocks.NewUserCacheMock(mc)
			return mock
		},

		userProducerMock: func(mc *minimock.Controller) kafka.Producer {
			mock := kafkaMocks.NewProducerMock(mc)
			return mock
		},
	}, {
		name: "user service: update user: empty user info case",
		args: args{
			ctx:      ctx,
			id:       id,
			userInfo: nil,
		},
		err: nilInfoErr,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			return mock
		},

		txRepositoryMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			return mock
		},

		userCacheMock: func(mc *minimock.Controller) repository.UserCache {
			mock := repoMocks.NewUserCacheMock(mc)
			return mock
		},

		userProducerMock: func(mc *minimock.Controller) kafka.Producer {
			mock := kafkaMocks.NewProducerMock(mc)
			return mock
		},
	}, {
		name: "user service: update user: user repo error case",
		args: args{
			ctx:      ctx,
			id:       id,
			userInfo: &userInfo,
		},
		err: userRepoErr,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.UpdateMock.Expect(ctx, id, userInfo).Return(userRepoErr)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			return mock
		},

		txRepositoryMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},

		userCacheMock: func(mc *minimock.Controller) repository.UserCache {
			mock := repoMocks.NewUserCacheMock(mc)
			return mock
		},

		userProducerMock: func(mc *minimock.Controller) kafka.Producer {
			mock := kafkaMocks.NewProducerMock(mc)
			return mock
		},
	}, {
		name: "user service: update user: log repo error case",
		args: args{
			ctx:      ctx,
			id:       id,
			userInfo: &userInfo,
		},
		err: logRepoErr,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.UpdateMock.Expect(ctx, id, userInfo).Return(nil)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			mock.SaveMock.Expect(ctx, logInfo).Return(logRepoErr)
			return mock
		},

		txRepositoryMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},

		userCacheMock: func(mc *minimock.Controller) repository.UserCache {
			mock := repoMocks.NewUserCacheMock(mc)
			return mock
		},

		userProducerMock: func(mc *minimock.Controller) kafka.Producer {
			mock := kafkaMocks.NewProducerMock(mc)
			return mock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			userRepoMock := tt.userRepositoryMock(mc)
			logRepoMock := tt.logRepositoyMock(mc)
			txManagerMock := tt.txRepositoryMock(mc)
			userCahceMock := tt.userCacheMock(mc)
			userProducerMock := tt.userProducerMock(mc)

			service := user.NewService(userRepoMock, logRepoMock, txManagerMock, userCahceMock, userProducerMock)
			err := service.UpdateUser(tt.args.ctx, tt.args.id, tt.args.userInfo)

			require.Equal(t, tt.err, err)
		})
	}
}

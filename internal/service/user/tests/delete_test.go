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

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoyMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txRepositoryMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache
	type userProducerMockFunc func(mc *minimock.Controller) kafka.Producer

	type args struct{}

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		id = gofakeit.Int64()

		logInfo = models.LogInfo{UserID: id, Action: consts.ActionDelete}

		saveUserErr = errors.New("save user error")
		saveLogErr  = errors.New("save log error")
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoyMock   logRepositoyMockFunc
		txRepositoryMock   txRepositoryMockFunc
		userCacheMock      userCacheMockFunc
		userProducerMock   userProducerMockFunc
	}{{
		name: "user service: delete user: success case",
		args: args{},
		err:  nil,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(nil)
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
		name: "user service: delete user: save user error case case",
		args: args{},
		err:  saveUserErr,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(saveUserErr)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			// mock.SaveMock.Expect(ctx, logInfo).Return(nil)
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
		name: "user service: delete user: write log error case",
		args: args{},
		err:  saveLogErr,

		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
			mock := repoMocks.NewUserRepositoryMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(nil)
			return mock
		},

		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			mock.SaveMock.Expect(ctx, logInfo).Return(saveLogErr)
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userRepoMock := tt.userRepositoryMock(mc)
			logRepoMock := tt.logRepositoyMock(mc)
			txManagerMock := tt.txRepositoryMock(mc)
			userCacheMock := tt.userCacheMock(mc)
			userProducerMock := tt.userProducerMock(mc)

			service := user.NewService(userRepoMock, logRepoMock, txManagerMock, userCacheMock, userProducerMock)
			err := service.DeleteUser(ctx, id)

			require.Equal(t, tt.err, err)
		})
	}
}

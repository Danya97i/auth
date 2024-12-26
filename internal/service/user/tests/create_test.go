package tests

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/Danya97i/platform_common/pkg/db"
	dbMocks "github.com/Danya97i/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Danya97i/auth/internal/client/kafka"
	kafkaMocks "github.com/Danya97i/auth/internal/client/kafka/mocks"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/repository"
	repoMocks "github.com/Danya97i/auth/internal/repository/mocks"
	"github.com/Danya97i/auth/internal/service/user"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoyMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txRepositoryMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache
	type userProducerMockFunc func(mc *minimock.Controller) kafka.Producer

	type args struct {
		ctx         context.Context
		userInfo    models.UserInfo
		pass        string
		passConfirm string
	}

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		pass         = gofakeit.Word()
		passConfirm  = pass
		wrongConfirm = gofakeit.Word()

		userInfo = models.UserInfo{
			Name:  &name,
			Email: email,
			Role:  consts.User,
		}

		providerMessage, _ = json.Marshal(&userInfo)

		logInfo = models.LogInfo{UserID: id, Action: consts.ActionCreate}

		ErrEmptyName     = errors.New("user name is empty")
		ErrWrongEmail    = errors.New("mail: no angle-addr")
		ErrPassword      = errors.New("passwords don't match")
		ErrUserRepoError = errors.New("repo error")
		ErrLogRepoError  = errors.New("log repo error")
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
	}{
		{
			name: "user service: create user: success case",
			args: args{
				ctx:         ctx,
				userInfo:    userInfo,
				pass:        pass,
				passConfirm: passConfirm,
			},
			want: id,
			err:  nil,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				rMock.CreateMock.Inspect(func(_ctx context.Context, _userInfo models.UserInfo, _ string) {
					assert.Equal(mc, _ctx, ctx)
					assert.Equal(mc, _userInfo, userInfo)
				}).Return(id, nil)
				return rMock
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
				mock.SendMessageMock.Expect(ctx, providerMessage).Return(nil)
				return mock
			},
		},

		{
			name: "user service: create user: empty name case",
			args: args{
				ctx: ctx,
				userInfo: models.UserInfo{
					Name:  nil,
					Email: email,
					Role:  consts.User,
				},
				pass:        pass,
				passConfirm: passConfirm,
			},
			want: 0,
			err:  ErrEmptyName,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				return rMock
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
		},

		{
			name: "user service: create user: incorrect email case",
			args: args{
				ctx: ctx,
				userInfo: models.UserInfo{
					Name:  &name,
					Email: "incorrect email",
					Role:  consts.User,
				},
				pass:        pass,
				passConfirm: passConfirm,
			},
			want: 0,
			err:  ErrWrongEmail,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				return rMock
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
		},

		{
			name: "user service: create user: wrong password confirm case",
			args: args{
				ctx: ctx,
				userInfo: models.UserInfo{
					Name:  &name,
					Email: email,
					Role:  consts.User,
				},
				pass:        pass,
				passConfirm: wrongConfirm,
			},
			want: 0,
			err:  ErrPassword,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				return rMock
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
		},

		{
			name: "user service: create user: user repo error case",
			args: args{
				ctx:         ctx,
				userInfo:    userInfo,
				pass:        pass,
				passConfirm: passConfirm,
			},
			want: 0,
			err:  ErrUserRepoError,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				rMock.CreateMock.Inspect(func(_ctx context.Context, _userInfo models.UserInfo, _ string) {
					assert.Equal(mc, _ctx, ctx)
					assert.Equal(mc, _userInfo, userInfo)
				}).Return(0, ErrUserRepoError)
				return rMock
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
		},

		{
			name: "user service: create user: log repo error case",
			args: args{
				ctx:         ctx,
				userInfo:    userInfo,
				pass:        pass,
				passConfirm: passConfirm,
			},
			want: 0,
			err:  ErrLogRepoError,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				rMock := repoMocks.NewUserRepositoryMock(mc)
				rMock.CreateMock.Inspect(func(_ctx context.Context, _userInfo models.UserInfo, _ string) {
					assert.Equal(mc, _ctx, ctx)
					assert.Equal(mc, _userInfo, userInfo)
				}).Return(id, nil)
				return rMock
			},

			logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repoMocks.NewLogRepositoryMock(mc)
				mock.SaveMock.Expect(ctx, logInfo).Return(ErrLogRepoError)
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name, "\n")
			userRepoMock := tt.userRepositoryMock(mc)
			logRepoMock := tt.logRepositoyMock(mc)
			txManagerMock := tt.txRepositoryMock(mc)
			userCacheMock := tt.userCacheMock(mc)
			userProducerMock := tt.userProducerMock(mc)

			service := user.NewService(userRepoMock, logRepoMock, txManagerMock, userCacheMock, userProducerMock)
			newID, err := service.CreateUser(tt.args.ctx, tt.args.userInfo, tt.args.pass, tt.args.passConfirm)

			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}

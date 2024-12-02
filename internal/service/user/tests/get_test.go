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

	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/repository"
	repoMocks "github.com/Danya97i/auth/internal/repository/mocks"
	"github.com/Danya97i/auth/internal/service/user"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoyMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txRepositoryMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache

	type args struct {
		сtx context.Context
		id  int64
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

		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		wantUser = &models.User{
			ID:        id,
			Info:      &userInfo,
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}

		logInfo = models.LogInfo{UserID: id, Action: consts.ActionGet}

		userRepoErr   = errors.New("get user error")
		logRepoErr    = errors.New("save log error")
		errNotInCache = errors.New("user not in cache")
	)

	tests := []struct {
		name               string
		args               args
		want               *models.User
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoyMock   logRepositoyMockFunc
		txRepositoryMock   txRepositoryMockFunc
		userCacheMock      userCacheMockFunc
	}{
		{
			name: "user service: get user: success repo case",
			args: args{
				сtx: ctx,
				id:  id,
			},
			want: wantUser,
			err:  nil,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UserMock.Expect(ctx, id).Return(wantUser, nil)
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
				mock.GetMock.Expect(ctx, id).Return(nil, errNotInCache)
				mock.SetMock.Expect(ctx, wantUser).Return(nil)
				return mock
			},
		},

		{
			name: "user service: get user: success cache case",
			args: args{
				сtx: ctx,
				id:  id,
			},
			want: wantUser,
			err:  nil,

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
				mock.GetMock.Expect(ctx, id).Return(wantUser, nil)
				return mock
			},
		},

		{
			name: "user service: get user: user repo error case",
			args: args{
				сtx: ctx,
				id:  id,
			},
			want: nil,
			err:  userRepoErr,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UserMock.Expect(ctx, id).Return(nil, userRepoErr)
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
				mock.GetMock.Expect(ctx, id).Return(nil, errNotInCache)
				return mock
			},
		},

		{
			name: "user service: get user: log repo error case",
			args: args{
				сtx: ctx,
				id:  id,
			},
			want: nil,
			err:  logRepoErr,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UserMock.Expect(ctx, id).Return(wantUser, nil)
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
				mock.GetMock.Expect(ctx, id).Return(nil, errNotInCache)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)
			userRepoMock := tt.userRepositoryMock(mc)
			logRepoMock := tt.logRepositoyMock(mc)
			txManagerMock := tt.txRepositoryMock(mc)
			userCahceMock := tt.userCacheMock(mc)

			service := user.NewService(userRepoMock, logRepoMock, txManagerMock, userCahceMock)
			gettedUser, err := service.User(tt.args.сtx, tt.args.id)
			require.Equal(t, tt.want, gettedUser)
			require.Equal(t, tt.err, err)
		})
	}
}

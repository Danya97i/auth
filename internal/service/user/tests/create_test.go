package tests

// import (
// 	"context"
// 	"testing"

// 	"github.com/Danya97i/platform_common/pkg/db"
// 	dbMocks "github.com/Danya97i/platform_common/pkg/db/mocks"
// 	"github.com/brianvoe/gofakeit"
// 	"github.com/gojuno/minimock/v3"
// 	"github.com/stretchr/testify/require"
// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/Danya97i/auth/internal/models"
// 	"github.com/Danya97i/auth/internal/models/consts"
// 	"github.com/Danya97i/auth/internal/repository"
// 	repoMocks "github.com/Danya97i/auth/internal/repository/mocks"
// 	"github.com/Danya97i/auth/internal/service/user"
// )

// func TestCreateUser(t *testing.T) {
// 	t.Parallel()

// 	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
// 	type logRepositoyMockFunc func(mc *minimock.Controller) repository.LogRepository
// 	type txRepositoryMockFunc func(mc *minimock.Controller) db.TxManager

// 	type args struct{}

// 	var (
// 		ctx = context.Background()

// 		mc = minimock.NewController(t)

// 		id    = gofakeit.Int64()
// 		name  = gofakeit.Name()
// 		email = gofakeit.Email()

// 		pass        = gofakeit.Word()
// 		passConfirm = pass
// 		// wrongConfirm = gofakeit.Word()

// 		passHash, _ = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

// 		userInfo = models.UserInfo{
// 			Name:  &name,
// 			Email: email,
// 			Role:  consts.User,
// 		}

// 		logInfo = models.LogInfo{UserID: id, Action: consts.ActionCreate}
// 	)

// 	tests := []struct {
// 		name               string
// 		args               args
// 		want               int64
// 		err                error
// 		userRepositoryMock userRepositoryMockFunc
// 		logRepositoyMock   logRepositoyMockFunc
// 		txRepositoryMock   txRepositoryMockFunc
// 	}{{
// 		name: "user service: create user: success case",
// 		args: args{},
// 		want: id,
// 		err:  nil,

// 		userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
// 			mock := repoMocks.NewUserRepositoryMock(mc)
// 			mock.CreateMock.Expect(ctx, userInfo, string(passHash)).Return(id, nil)
// 			return mock
// 		},

// 		logRepositoyMock: func(mc *minimock.Controller) repository.LogRepository {
// 			mock := repoMocks.NewLogRepositoryMock(mc)
// 			mock.SaveMock.Expect(ctx, logInfo).Return(nil)
// 			return mock
// 		},

// 		txRepositoryMock: func(mc *minimock.Controller) db.TxManager {
// 			mock := dbMocks.NewTxManagerMock(mc)
// 			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
// 				return f(ctx)
// 			})
// 			return mock
// 		},
// 	}}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()
// 			userRepoMock := tt.userRepositoryMock(mc)
// 			logRepoMock := tt.logRepositoyMock(mc)
// 			txManagerMock := tt.txRepositoryMock(mc)

// 			service := user.NewService(userRepoMock, logRepoMock, txManagerMock)
// 			newID, err := service.CreateUser(ctx, userInfo, pass, passConfirm)

// 			require.Equal(t, tt.want, newID)
// 			require.Equal(t, tt.err, err)
// 		})
// 	}
// }

package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/service"
	serviceMock "github.com/Danya97i/auth/internal/service/mocks"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

func TestCreateUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *pb.CreateUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Word()

		req = &pb.CreateUserRequest{
			Info: &pb.UserInfo{
				Name:  name,
				Email: email,
				Role:  pb.Role_USER,
			},
			Password:        password,
			PasswordConfirm: password,
		}

		resp = &pb.CreateUserResponse{
			Id: id,
		}

		userInfo = models.UserInfo{
			Name:  &name,
			Email: email,
			Role:  consts.User,
		}

		emptyUserInfoErr = errors.New("invalid info")
		serviceErr       = errors.New("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *pb.CreateUserResponse
		err             error
		userServiceMock userServiceMockFunc
	}{{
		name: "server: create user: success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: resp,
		err:  nil,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.CreateUserMock.Expect(ctx, userInfo, password, password).Return(id, nil)
			return userServiceMock
		},
	}, {
		name: "server: create user: nil user info case",
		args: args{
			ctx: ctx,
			req: &pb.CreateUserRequest{},
		},
		want: nil,
		err:  emptyUserInfoErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			return userServiceMock
		},
	}, {
		name: "server: create user: service error case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  serviceErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.CreateUserMock.Expect(ctx, userInfo, password, password).Return(0, serviceErr)
			return userServiceMock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewServer(userServiceMock)
			newID, err := api.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}

package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/service"
	serviceMock "github.com/Danya97i/auth/internal/service/mocks"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

func TestGetUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *pb.GetUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		req = &pb.GetUserRequest{
			Id: id,
		}

		resp = &pb.GetUserResponse{
			User: &pb.User{
				Id: id,
				Info: &pb.UserInfo{
					Name:  name,
					Email: email,
					Role:  pb.Role_USER,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}

		userInfo = models.UserInfo{
			Name:  &name,
			Email: email,
			Role:  consts.User,
		}

		userModel = models.User{
			ID:        id,
			Info:      &userInfo,
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}

		serviceErr = errors.New("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *pb.GetUserResponse
		err             error
		userServiceMock userServiceMockFunc
	}{{
		name: "server: get user: success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: resp,
		err:  nil,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.UserMock.Expect(ctx, id).Return(&userModel, nil)
			return userServiceMock
		},
	}, {
		name: "server: get user: service error case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  serviceErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.UserMock.Expect(ctx, id).Return(nil, serviceErr)
			return userServiceMock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewServer(userServiceMock)
			gettedUser, err := api.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, gettedUser)
			require.Equal(t, tt.err, err)
		})
	}
}

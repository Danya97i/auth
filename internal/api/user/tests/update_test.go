package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/models"
	"github.com/Danya97i/auth/internal/models/consts"
	"github.com/Danya97i/auth/internal/service"
	serviceMock "github.com/Danya97i/auth/internal/service/mocks"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

func TestUpdateUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *pb.UpdateUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Name()

		req = &pb.UpdateUserRequest{
			Id:   id,
			Name: wrapperspb.String(name),
			Role: pb.Role_ADMIN,
		}

		resp *emptypb.Empty = nil

		userInfo = &models.UserInfo{
			Name: &name,
			Role: consts.Admin,
		}

		serviceErr = errors.New("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{{
		name: "server: update user: success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: resp,
		err:  nil,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.UpdateUserMock.Expect(ctx, id, userInfo).Return(nil)
			return userServiceMock
		},
	}, {
		name: "server: update user: service error case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  serviceErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.UpdateUserMock.Expect(ctx, id, userInfo).Return(serviceErr)
			return userServiceMock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewServer(userServiceMock)
			newID, err := api.UpdateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}

package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Danya97i/auth/internal/api/user"
	"github.com/Danya97i/auth/internal/service"
	serviceMock "github.com/Danya97i/auth/internal/service/mocks"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

func TestDeleteUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *pb.DeleteUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		req = &pb.DeleteUserRequest{
			Id: id,
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
		name: "server: delete user: success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  nil,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.DeleteUserMock.Expect(ctx, id).Return(nil)
			return userServiceMock
		},
	}, {
		name: "server: delete user: service error case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  serviceErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			userServiceMock := serviceMock.NewUserServiceMock(mc)
			userServiceMock.DeleteUserMock.Expect(ctx, id).Return(serviceErr)
			return userServiceMock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewServer(userServiceMock)
			_, err := api.DeleteUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}

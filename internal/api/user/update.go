package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Danya97i/auth/internal/converter"
	"github.com/Danya97i/auth/internal/models"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// UpdateUser - метод обновления данных пользователя
func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Println("update user request: ", req)

	userInfo := models.UserInfo{
		Role: converter.ToRoleFromPb(req.Role),
	}

	if req.Name != nil {
		userInfo.Name = &req.Name.Value
	}

	if err := s.userService.UpdateUser(ctx, req.Id, &userInfo); err != nil {
		return nil, err
	}

	return nil, nil
}
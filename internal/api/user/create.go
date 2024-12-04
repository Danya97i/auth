package user

import (
	"context"
	"errors"
	"log"

	"github.com/Danya97i/auth/internal/converter"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// CreateUser - метод для создания нового пользователя
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("create user request: ", req)

	info := converter.ToUserInfoFromPb(req.Info)
	if info == nil {
		return nil, errors.New("invalid info")
	}

	id, err := s.userService.CreateUser(ctx, *info, req.Password, req.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Id: id}, nil
}

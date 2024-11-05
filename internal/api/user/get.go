package user

import (
	"context"
	"log"

	"github.com/Danya97i/auth/internal/converter"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// GetUser - метод для получения информации о пользователе
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("get user request: ", req)

	user, err := s.userService.User(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: converter.ToUserFromService(user),
	}, nil
}

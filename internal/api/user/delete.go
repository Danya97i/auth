package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// DeleteUser - метод удаления пользователя
func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Println("delete user request: ", req)

	if err := s.userService.DeleteUser(ctx, req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

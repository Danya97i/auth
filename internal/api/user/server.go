package user

import (
	"github.com/Danya97i/auth/internal/service"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// Server - имплементация gRPC-сервера
type Server struct {
	pb.UnimplementedUserV1Server
	userService service.UserService
}

// NewServer - конструктор сервера
func NewServer(userService service.UserService) *Server {
	return &Server{userService: userService}
}

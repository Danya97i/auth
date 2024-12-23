package auth

import (
	"github.com/Danya97i/auth/internal/service"
	pb "github.com/Danya97i/auth/pkg/auth_v1"
)

// Server is the server for the auth service
type Server struct {
	pb.UnimplementedAuthV1Server
	authService service.AuthService
}

// NewServer creates a new server
func NewServer(authService service.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}

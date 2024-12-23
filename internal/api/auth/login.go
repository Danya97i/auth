package auth

import (
	"context"

	pb "github.com/Danya97i/auth/pkg/auth_v1"
)

// Login handles the login request
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	refreshToken, err := s.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}

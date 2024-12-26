package auth

import (
	"context"

	pb "github.com/Danya97i/auth/pkg/auth_v1"
)

// GetRefreshToken returns a refresh token
func (s *Server) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error) {
	token, err := s.authService.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.GetRefreshTokenResponse{
		RefreshToken: token,
	}, nil
}

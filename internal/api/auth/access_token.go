package auth

import (
	"context"

	pb "github.com/Danya97i/auth/pkg/auth_v1"
)

// GetAccessToken returns a access token
func (s *Server) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error) {
	token, err := s.authService.GetAccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccessTokenResponse{
		AccessToken: token,
	}, nil
}

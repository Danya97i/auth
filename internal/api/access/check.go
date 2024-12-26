package access

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/auth/pkg/access_v1"
)

const (
	authHeaderName = "authorization"
	authPrefix     = "Bearer "
)

// Check check user access to endpoint
func (s *Server) Check(ctx context.Context, req *pb.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md[authHeaderName]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("authorization header is not valid")
	}

	token := strings.TrimPrefix(authHeader[0], authPrefix)

	if err := s.accessService.Check(ctx, token, req.EndpointAddress); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

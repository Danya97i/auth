package access

import (
	"github.com/Danya97i/auth/internal/service"
	pb "github.com/Danya97i/auth/pkg/access_v1"
)

// Server – access server
type Server struct {
	pb.UnimplementedAccessV1Server
	accessService service.AccessService
}

// NewServer – creates new access server
func NewServer(
	accessService service.AccessService,
) *Server {
	return &Server{
		accessService: accessService,
	}
}

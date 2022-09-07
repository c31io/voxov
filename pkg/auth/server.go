package auth

import (
	pb "github.com/c31io/voxov/pkg/api/auth"
	"github.com/c31io/voxov/pkg/health"
)

var Health *health.Health

type Server struct {
	pb.UnimplementedAuthServer
}

func init() {
	Health = &health.Health{
		Alive: true,
		Dead:  make(chan struct{}),
	}
}

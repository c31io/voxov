package gate

import (
	pb "github.com/c31io/voxov/api"
)

var (
	_ = "env"
)

type Server struct {
	pb.UnimplementedVOxOVServer
}

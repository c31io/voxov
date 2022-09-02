package auth

import (
	pb "github.com/c31io/voxov/pkg/api/auth"
)

type Server struct {
	pb.UnimplementedAuthServer
}

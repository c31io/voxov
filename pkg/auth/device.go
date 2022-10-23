package auth

import (
	"context"

	pb "github.com/c31io/voxov/pkg/api/auth"
)

func (s *Server) NewDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	return &pb.Device{}, nil
}

package gate

import (
	"context"
	"log"

	pb "github.com/c31io/voxov/api"
	pbAuth "github.com/c31io/voxov/pkg/api/auth"
)

// Create a new device
func (s *Server) CreateDevice(ctx context.Context, in *pb.CreateDeviceRequest) (*pb.CreateDeviceReply, error) {
	if !isValidToken(ctx, in.GetToken()) {
		return &pb.CreateDeviceReply{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.NewDevice(ctx, &pbAuth.NewDeviceRequest{Token: in.GetToken(), Dtoken: in.GetDtoken(), Name: in.GetName(), Info: in.GetInfo()})
	if err != nil {
		log.Println("Failed to get NewDeviceReply")
		Health.NowDead()
		return &pb.CreateDeviceReply{}, nil
	}
	return &pb.CreateDeviceReply{Ok: r.GetOk()}, nil
}

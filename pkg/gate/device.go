package gate

import (
	"context"
	"log"

	pb "github.com/c31io/voxov/api"
	pbAuth "github.com/c31io/voxov/pkg/api/auth"
)

// Create a new device
func (s *Server) CreateDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	if !isValidToken(ctx, in.GetToken()) {
		return &pb.Device{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.NewDevice(ctx, &pbAuth.Device{
		Dname: in.GetDname(),
		Dinfo: in.GetDinfo(),
	})
	if err != nil {
		log.Println("Failed to get NewDeviceReply")
		Health.NowDead()
		return &pb.Device{}, nil
	}
	return &pb.Device{
		Did:     r.GetDid(),
		Dtoken:  r.GetDtoken(),
		Dname:   r.GetDname(),
		Dinfo:   r.GetDinfo(),
		Pid:     r.GetPid(),
		Created: r.GetCreated(),
		LastIn:  r.GetLastIn(),
	}, nil
}

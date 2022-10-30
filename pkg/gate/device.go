package gate

import (
	"context"
	"log"

	pb "github.com/c31io/voxov/api"
	pbAuth "github.com/c31io/voxov/pkg/api/auth"
	r "github.com/c31io/voxov/pkg/rdb"
	"github.com/go-redis/redis/v9"
)

// Create a new device
func (s *Server) CreateDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	pid := getPid(ctx, in.GetToken())
	if pid == 0 {
		log.Println("Failed to get pid")
		return &pb.Device{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.NewDevice(ctx, &pbAuth.Device{
		Dname: in.GetDname(),
		Dinfo: in.GetDinfo(),
		Pid:   pid,
	})
	if err != nil {
		log.Println("Failed to get NewDeviceReply")
		Health.NowDead()
		return &pb.Device{}, nil
	}
	log.Println("Created device")
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

func (s *Server) ReadDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	pid := getPid(ctx, in.GetToken())
	if pid == 0 {
		log.Println("Failed to get pid")
		return &pb.Device{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.GetDevice(ctx, &pbAuth.Device{
		Did: in.GetDid(),
		Pid: pid,
	})
	if err != nil {
		log.Println("Failed to read device: " + err.Error())
		return &pb.Device{}, nil
	}
	log.Println("Read device")
	return &pb.Device{
		Did:     r.GetDid(),
		Dname:   r.GetDname(),
		Dinfo:   r.GetDinfo(),
		Pid:     r.GetPid(),
		Created: r.GetCreated(),
		LastIn:  r.GetLastIn(),
	}, nil
}

func (s *Server) UpdateDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	pid := getPid(ctx, in.GetToken())
	if pid == 0 {
		log.Println("Failed to get pid")
		return &pb.Device{}, nil
	}
	// TODO impl
	return &pb.Device{}, nil
}

func (s *Server) AuthDevice(ctx context.Context, in *pb.AuthDeviceRequest) (*pb.AuthDeviceReply, error) {
	if !isValidToken(ctx, in.GetToken()) {
		log.Println("Invalid token")
		return &pb.AuthDeviceReply{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	reply, err := c.CheckDevice(ctx, &pbAuth.CheckDeviceRequest{
		Dtoken: in.GetDtoken(),
	})
	if err != nil {
		log.Println("Failed to AuthDevice: " + err.Error())
		Health.NowDead()
		return &pb.AuthDeviceReply{}, nil
	}
	err = rdb.SetXX(ctx, "s"+string(in.GetToken()), r.Int64ToByteSlice(reply.GetPid()), redis.KeepTTL).Err()
	if err == redis.Nil {
		log.Println("Session expired during authentication")
		return &pb.AuthDeviceReply{}, nil
	} else if err != nil {
		log.Println("Failed to set rdb: " + err.Error())
		Health.NowDead()
		return &pb.AuthDeviceReply{}, nil
	}
	log.Println("AuthDevice")
	return &pb.AuthDeviceReply{
		Did: reply.GetDid(),
		Pid: reply.GetPid(),
	}, nil
}

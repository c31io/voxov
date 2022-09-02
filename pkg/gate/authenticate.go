package gate

import (
	"context"
	"log"
	"os"
	"strconv"

	pb "github.com/c31io/voxov/api"
)

var authMaxTtl int64 // seconds

func init() {
	val, ok := os.LookupEnv("AUTH_MAX_TTL")
	if !ok {
		authMaxTtl = 3600
	} else {
		i, err := strconv.Atoi(val)
		if err != nil || i <= 0 {
			log.Fatal("AUTH_MAX_TTL must be a positive integer")
		}
		authMaxTtl = int64(i)
	}
}

func (s *Server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	return &pb.AuthenticateReply{}, nil
}

func (s *Server) WhoAmI(ctx context.Context, in *pb.WhoAmIRequest) (*pb.WhoAmIReply, error) {
	return &pb.WhoAmIReply{}, nil
}

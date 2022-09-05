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

// Client requests what to send to who to authenticate.
func (s *Server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	if !isValidToken(ctx, &in.Token) || in.Ttl <= 0 {
		return &pb.AuthenticateReply{}, nil
	}
	ttl := minInt64(authMaxTtl, in.GetTtl())
	_ = ttl
	// TODO: pkg auth
	return &pb.AuthenticateReply{}, nil
}

// Client checks if sms was received, and if so update session person.
func (s *Server) WhoAmI(ctx context.Context, in *pb.WhoAmIRequest) (*pb.WhoAmIReply, error) {
	return &pb.WhoAmIReply{}, nil
}

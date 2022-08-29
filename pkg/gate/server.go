package gate

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/c31io/voxov/api"
)

type Server struct {
	pb.UnimplementedVOxOVServer
}

var apiVersion int32

func init() {
	val, ok := os.LookupEnv("API_VERSION")
	if !ok {
		apiVersion = 1
	} else {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("API_VERSION should be an integer")
		}
		apiVersion = int32(i)
	}
}

func (s *Server) CreateSession(ctx context.Context, in *pb.CreateSessionRequest) (*pb.CreateSessionReply, error) {
	if apiVersion != in.GetApiVersion() || in.GetTtl() <= 0 {
		log.Printf("Client API version: %d", in.GetApiVersion())
		return &pb.CreateSessionReply{ApiVersion: apiVersion}, nil
	}
	token := genToken()
	err := rdb.Set(ctx, string(token), "", time.Duration(in.GetTtl())*time.Second).Err()
	if err != nil {
		log.Print("Failed to set on rdb")
		return &pb.CreateSessionReply{ApiVersion: apiVersion}, nil
	}
	return &pb.CreateSessionReply{ApiVersion: apiVersion, Token: token}, nil
}

func (s *Server) UpdateSession(ctx context.Context, in *pb.UpdateSessionRequest) (*pb.UpdateSessionReply, error) {
	if in.GetTtl() <= 0 {
		return &pb.UpdateSessionReply{Ok: false}, nil
	}
	err := rdb.ExpireXX(ctx, string(in.GetToken()), time.Duration(in.GetTtl())*time.Second).Err()
	if err != nil {
		return &pb.UpdateSessionReply{Ok: false}, nil
	}
	return &pb.UpdateSessionReply{Ok: true}, nil
}

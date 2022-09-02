package gate

import (
	"context"
	"crypto/rand"
	"log"
	"time"

	pb "github.com/c31io/voxov/api"
)

// It takes more than 1.7*10^8 years to randomly guess one
// in 10^12 valid sessions with a rate of 10^18Hz
const tokenBytes = 16

func genToken() []byte {
	b := make([]byte, tokenBytes)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("crypto/rand returned an error")
	}
	return b
}

// Initiate a session with a positive ttl.
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

// Assign ttl if session is not expired.
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

func isValidToken(ctx context.Context, token *[]byte) bool {
	val, err := rdb.Exists(ctx, string(*token)).Result()
	if err != nil || val == 0 {
		return false
	} else {
		return true
	}
}

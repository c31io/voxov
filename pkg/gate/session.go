package gate

import (
	"context"
	"crypto/rand"
	"log"
	"os"
	"strconv"
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

var sessionMaxTtl int64 // seconds

func init() {
	val, ok := os.LookupEnv("SESSION_MAX_TTL")
	if !ok {
		sessionMaxTtl = 3600
	} else {
		i, err := strconv.Atoi(val)
		if err != nil || i <= 0 {
			log.Fatal("SESSION_MAX_TTL must be a positive integer")
		}
		sessionMaxTtl = int64(i)
	}
}

// Initiate a session with a positive ttl.
func (s *Server) CreateSession(ctx context.Context, in *pb.CreateSessionRequest) (*pb.CreateSessionReply, error) {
	if apiVersion != in.GetApiVersion() || in.GetTtl() <= 0 {
		log.Printf("Client API version: %d", in.GetApiVersion())
		return &pb.CreateSessionReply{ApiVersion: apiVersion}, nil
	}
	token := genToken()
	ttl := minInt64(sessionMaxTtl, in.GetTtl())
	err := rdb.Set(ctx, "s"+string(token), "", time.Duration(ttl)*time.Second).Err()
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
	ttl := minInt64(sessionMaxTtl, in.GetTtl())
	err := rdb.ExpireXX(ctx, "s"+string(in.GetToken()), time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return &pb.UpdateSessionReply{Ok: false}, nil
	}
	return &pb.UpdateSessionReply{Ok: true}, nil
}

func isValidToken(ctx context.Context, token *[]byte) bool {
	val, err := rdb.Exists(ctx, "s"+string(*token)).Result()
	if err != nil || val == 0 {
		return false
	} else {
		return true
	}
}

// go:build !aws

// sms mock
package auth

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/c31io/voxov/pkg/api/auth"
)

var waitMsgTtl time.Duration

func init() {
	val, ok := os.LookupEnv("WAIT_MSG_TTL")
	if !ok {
		waitMsgTtl = 10 * time.Minute
	} else {
		i, err := strconv.Atoi(val)
		if err != nil || i <= 0 {
			log.Fatal("WAIT_MSG_TTL must be a positive integer")
		}
		waitMsgTtl = time.Duration(i) * time.Second
	}
}

func (s *Server) WaitMsg(ctx context.Context, in *pb.WaitMsgRequest) (*pb.WaitMsgReply, error) {
	return &pb.WaitMsgReply{}, nil
}

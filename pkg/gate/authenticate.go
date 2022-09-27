package gate

import (
	"context"
	"log"
	"os"

	pb "github.com/c31io/voxov/api"
	pbAuth "github.com/c31io/voxov/pkg/api/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authAddr string
	authConn *grpc.ClientConn
)

func init() {
	val, ok := os.LookupEnv("AUTH_ADDR")
	if !ok {
		authAddr = "localhost:10002"
	} else {
		authAddr = val
	}
	var err error
	authConn, err = grpc.Dial(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
}

// Client requests what to send to who to authenticate.
func (s *Server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	if !isValidToken(ctx, in.GetToken()) {
		return &pb.AuthenticateReply{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.WaitMsg(ctx, &pbAuth.WaitMsgRequest{Token: in.GetToken()})
	if err != nil {
		log.Println("Failed to get WaitMsgReply")
		Health.NowDead()
		return &pb.AuthenticateReply{}, nil
	}
	return &pb.AuthenticateReply{Tel: r.GetTel(), Msg: r.GetMsg()}, nil
}

// Client checks if sms was received, and if so update session person.
func (s *Server) WhoAmI(ctx context.Context, in *pb.WhoAmIRequest) (*pb.WhoAmIReply, error) {
	if !isValidToken(ctx, in.GetToken()) {
		return &pb.WhoAmIReply{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.CheckMsg(ctx, &pbAuth.CheckMsgRequest{Token: in.GetToken(), Tel: in.GetTel(), Msg: in.GetMsg()})
	if err != nil {
		log.Println("Failed to get WhoAmIReply")
		Health.NowDead()
		return &pb.WhoAmIReply{}, nil
	}
	return &pb.WhoAmIReply{Person: r.GetPerson()}, nil
}

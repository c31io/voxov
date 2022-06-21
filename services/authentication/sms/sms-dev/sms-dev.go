package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/c31io/voxov/services/authentication/sms/api"
	sms "github.com/c31io/voxov/services/authentication/sms/pkg"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50001, "The server port")
)

type server struct {
	pb.UnimplementedVeriCodeServer
}

func (s *server) SendVeriCode(ctx context.Context, in *pb.VeriCodeRequest) (*pb.VeriCodeReply, error) {
	log.Printf("Phone number: %v-%v", in.GetCoutryCode(), in.GetSubscriberNumber())
	code, err := sms.GenCode()
	if err != nil {
		return &pb.VeriCodeReply{VeriCode: uint32(code), Error: 1}, err
	}
	// SMS vendor-specific code goes here
	log.Printf("%v-%v got %d", in.GetCoutryCode(), in.GetSubscriberNumber(), code)
	return &pb.VeriCodeReply{VeriCode: uint32(code)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVeriCodeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

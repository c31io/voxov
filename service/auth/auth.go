package main

import (
	"log"
	"net"
	"os"

	"github.com/c31io/voxov/pkg/auth"
	"google.golang.org/grpc"

	pb "github.com/c31io/voxov/pkg/api/auth"
)

var authPort string

func init() {
	val, ok := os.LookupEnv("AUTH_PORT")
	if !ok {
		authPort = "10002"
	} else {
		authPort = val
	}
}

func main() {
	lis, err := net.Listen("tcp", ":"+authPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &auth.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

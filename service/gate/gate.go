package main

import (
	"log"
	"net"
	"os"

	"github.com/c31io/voxov/pkg/gate"
	"google.golang.org/grpc"

	pb "github.com/c31io/voxov/api"
)

var gatePort string

func init() {
	val, ok := os.LookupEnv("GATE_PORT")
	if !ok {
		gatePort = "10001"
	} else {
		gatePort = val
	}
}

func main() {
	lis, err := net.Listen("tcp", gatePort)
	if err != nil {
		log.Fatalf("gate failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVOxOVServer(s, &gate.Server{})
	log.Printf("gate server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gate failed to serve: %v", err)
	}
}

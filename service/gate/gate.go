package main

import (
	"log"
	"net"
	"os"

	"github.com/c31io/voxov/pkg/gate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

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
	lis, err := net.Listen("tcp", ":"+gatePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVOxOVServer(s, &gate.Server{})
	grpc_health_v1.RegisterHealthServer(s, gate.Health)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

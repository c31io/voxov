package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/c31io/voxov/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr       = flag.String("addr", "localhost:10001", "the address to connect to")
	apiVersion = 1
	ttl        = 3600
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVOxOVClient(conn)
	ctx := context.Background()
	r, err := c.CreateSession(ctx, &pb.CreateSessionRequest{ApiVersion: int32(apiVersion), Ttl: int64(ttl)})
	if err != nil {
		log.Fatal("Failed to create session")
	}
	log.Printf("Server API version: %d\n", r.GetApiVersion())
	token := r.GetToken()
	log.Printf("Session token: %s\n", hex.EncodeToString(token))
	go func() {
		for {
			time.Sleep(time.Duration(ttl/2) * time.Second)
			r, err := c.UpdateSession(ctx, &pb.UpdateSessionRequest{Token: token, Ttl: int64(ttl)})
			if !r.GetOk() || err != nil {
				log.Fatal("Failed to update session")
			}
			log.Println("Session updated")
		}
	}()
	for {
		s := new(string)
		fmt.Print("> ")
		fmt.Scanln(s)
	}
}

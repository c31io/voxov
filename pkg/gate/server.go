package gate

import (
	"log"
	"os"
	"strconv"

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

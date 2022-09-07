package gate

import (
	"log"
	"os"
	"strconv"

	pb "github.com/c31io/voxov/api"
	"github.com/c31io/voxov/pkg/health"
)

type Server struct {
	pb.UnimplementedVOxOVServer
}

var (
	apiVersion int32
	Health     *health.Health
)

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
	Health = &health.Health{
		Alive: true,
		Dead:  make(chan struct{}),
	}
}

package auth

import (
	"log"

	r "github.com/c31io/voxov/pkg/rdb"
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func init() {
	if r.Rdb != nil {
		rdb = r.Rdb
	} else {
		log.Fatal("rdb init failed")
	}
}

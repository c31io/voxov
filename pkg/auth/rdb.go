package auth

import (
	"os"

	"github.com/go-redis/redis/v9"
)

var (
	rdbAddr   string
	rdbPasswd string
	rdb       *redis.Client
)

func init() {
	val, ok := os.LookupEnv("RDB_ADDR")
	if !ok {
		rdbAddr = "localhost:6379"
	} else {
		rdbAddr = val
	}
	val, ok = os.LookupEnv("RDB_PASSWD")
	if !ok {
		rdbPasswd = ""
	} else {
		rdbPasswd = val
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     rdbAddr,
		Password: rdbPasswd,
	})
}

func minInt64(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

package pdb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	connStr string
	Pdb     *sql.DB
)

func init() {
	val, ok := os.LookupEnv("CONN_STR")
	if !ok {
		connStr = "user=voxov password=forLocalDevOnly dbname=pdb sslmode=disable"
	} else {
		connStr = val
	}
	var err error
	Pdb, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

package auth

import (
	"database/sql"
	"log"

	p "github.com/c31io/voxov/pkg/pdb"
)

var pdb *sql.DB

func init() {
	if p.Pdb != nil {
		pdb = p.Pdb
	} else {
		log.Fatal("rdb init failed")
	}
}

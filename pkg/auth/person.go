package auth

import (
	"context"
	"log"
	"time"

	pb "github.com/c31io/voxov/pkg/api/auth"
	p "github.com/c31io/voxov/pkg/pdb"
)

func (s *Server) GetPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	pid := in.GetPid()
	row := pdb.QueryRowContext(ctx, `SELECT (pid, hid, balance, phone, pname, id_doc, dlimit, created, last_in) FROM people
	WHERE pid = $1`, pid)
	var hid, balance int64
	var phone, pname, id_doc string
	var dlimit int32
	var created, last_in time.Time
	err := row.Scan(pid, hid, balance, phone, pname, id_doc, dlimit, created, last_in)
	if err != nil {
		log.Println("error GetPerson")
		return &pb.Person{}, nil
	}
	log.Println("GetPerson")
	return &pb.Person{
		Pid:     pid,
		Hid:     hid,
		Balance: balance,
		Phone:   phone,
		Pname:   pname,
		IdDoc:   id_doc,
		Dlimit:  dlimit,
		Created: p.TimeToUs(created),
		LastIn:  p.TimeToUs(last_in),
	}, nil
}

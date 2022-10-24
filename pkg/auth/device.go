package auth

import (
	"context"
	"crypto/rand"
	"log"
	"time"

	pb "github.com/c31io/voxov/pkg/api/auth"
	p "github.com/c31io/voxov/pkg/pdb"
)

// This is even longer than that is in gate/session.go
const tokenBytes = 24

func genToken() []byte {
	b := make([]byte, tokenBytes)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("crypto/rand returned an error")
	}
	return b
}

func (s *Server) NewDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	dtoken := genToken()
	dname := in.GetDname()
	dinfo := in.GetDinfo()
	pid := in.GetPid()
	rows, err := pdb.QueryContext(ctx, `INSERT INTO devices (dtoken, dname, dinfo, pid, created, last_in)
	VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp)
	ON CONFLICT (dtoken) DO NOTHING;
	RETURNING (did, created, last_in)`, dtoken, dname, dinfo, pid)
	if err != nil {
		log.Println("Failed to create new device: " + err.Error())
		Health.NowDead()
		return &pb.Device{}, nil
	}
	var did int64
	var created, last_in time.Time
	err = rows.Scan(did, created, last_in)
	if err != nil {
		log.Println("Failed to scan new device:" + err.Error())
		Health.NowDead()
		return &pb.Device{}, nil
	}
	return &pb.Device{
		Did:     did,
		Dtoken:  dtoken,
		Dname:   dname,
		Dinfo:   dinfo,
		Pid:     pid,
		Created: p.TimeToMs(created),
		LastIn:  p.TimeToMs(last_in),
	}, nil
}

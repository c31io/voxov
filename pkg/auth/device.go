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
	// TODO check dlimit, if exceeded, update the coldest device
	row := pdb.QueryRowContext(ctx, `INSERT INTO devices (dtoken, dname, dinfo, pid, created, last_in)
	VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp)
	ON CONFLICT (dtoken) DO NOTHING;
	RETURNING (did, created, last_in)`, dtoken, dname, dinfo, pid)
	var did int64
	var created, last_in time.Time
	err := row.Scan(did, created, last_in)
	if err != nil {
		log.Println("Failed to scan new device:" + err.Error())
		Health.NowDead()
		return &pb.Device{}, nil
	}
	log.Println("Create device")
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

func (s *Server) GetDevice(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	did := in.GetDid()
	pid := in.GetPid()
	row := pdb.QueryRowContext(ctx, `SELECT (did, dname, dinfo, pid, created, last_in) FROM devices
	WHERE did = $1 AND pid = $2`, did, pid)
	var dname, dinfo string
	var created, last_in time.Time
	err := row.Scan(did, dname, dinfo, pid, created, last_in)
	if err != nil {
		log.Println("Device not found")
		return &pb.Device{}, nil
	}
	log.Println("Got device")
	return &pb.Device{
		Did:     did,
		Dname:   dname,
		Dinfo:   in.Dinfo,
		Pid:     pid,
		Created: p.TimeToMs(created),
		LastIn:  p.TimeToMs(last_in),
	}, nil
}

func (s *Server) CheckDevice(ctx context.Context, in *pb.CheckDeviceRequest) (*pb.CheckDeviceReply, error) {
	dtoken := in.GetDtoken()
	row := pdb.QueryRowContext(ctx, `SELECT did FROM devices
	WHERE dtoken = $1`, dtoken)
	var did int64
	err := row.Scan(did)
	if err != nil {
		log.Println("dtoken not found")
		return &pb.CheckDeviceReply{}, nil
	}
	log.Println("CheckDevice")
	return &pb.CheckDeviceReply{
		Did: did,
	}, nil
}

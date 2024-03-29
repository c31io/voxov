package gate

import (
	"context"
	"log"

	pb "github.com/c31io/voxov/api"
	pbAuth "github.com/c31io/voxov/pkg/api/auth"
)

func (s *Server) ReadPerson(ctx context.Context, in *pb.Person) (*pb.Person, error) {
	if getPid(ctx, in.GetToken()) != in.GetPid() {
		log.Println("Token pid mismatch")
		return &pb.Person{}, nil
	}
	c := pbAuth.NewAuthClient(authConn)
	r, err := c.GetPerson(ctx, &pbAuth.Person{
		Pid: in.GetPid(),
	})
	if err != nil {
		log.Println("Failed to GetPerson")
		Health.NowDead()
		return &pb.Person{}, nil
	}
	log.Println("ReadPerson")
	return &pb.Person{
		Pid:     r.GetPid(),
		Hid:     r.GetHid(),
		Balance: r.GetBalance(),
		Phone:   r.GetPhone(),
		Pname:   r.GetPname(),
		IdDoc:   r.GetIdDoc(),
		Dlimit:  r.GetDlimit(),
		Created: r.GetCreated(),
		LastIn:  r.GetLastIn(),
	}, nil
}

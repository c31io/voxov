// go:build !aws

// sms mock
package auth

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	pb "github.com/c31io/voxov/pkg/api/auth"
	r "github.com/c31io/voxov/pkg/rdb"
	"github.com/go-redis/redis/v9"
)

var (
	waitMsgTtl time.Duration
	tels       []string
)

func init() {
	val, ok := os.LookupEnv("WAIT_MSG_TTL")
	if !ok {
		waitMsgTtl = 10 * time.Minute
	} else {
		i, err := strconv.Atoi(val)
		if err != nil || i <= 0 {
			log.Fatal("WAIT_MSG_TTL must be a positive integer")
		}
		waitMsgTtl = time.Duration(i) * time.Second
	}
	val, ok = os.LookupEnv("TELS")
	if !ok {
		log.Fatal("TELS must not be empty")
	} else {
		tels = strings.Split(val, ":")
		re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		for _, tel := range tels {
			if !re.MatchString(tel) {
				log.Fatal("Wrong TELS format")
			}
		}
	}
}

func (s *Server) WaitMsg(ctx context.Context, in *pb.WaitMsgRequest) (*pb.WaitMsgReply, error) {
	tel := tels[rand.Intn(len(tels))]
	code, err := GenCode()
	msg := fmt.Sprint(code)
	if err != nil {
		log.Fatal("crypto/rand returned an error")
	}
	val, err := rdb.SetNX(ctx, "m"+tel+msg, string(in.GetToken()), waitMsgTtl).Result()
	if err != nil {
		log.Println("Failed to setnx on rdb")
		Health.NowDead()
		return &pb.WaitMsgReply{}, nil
	}
	if !val {
		log.Println("The key was not set")
		return &pb.WaitMsgReply{}, nil
	}
	return &pb.WaitMsgReply{Tel: tel, Msg: msg}, nil
}

// Received? New person?
func (s *Server) CheckMsg(ctx context.Context, in *pb.CheckMsgRequest) (*pb.CheckMsgReply, error) {
	tel := in.GetTel()
	msg := in.GetMsg()
	token, err := rdb.Get(ctx, "m"+tel+msg).Result()
	if err == redis.Nil {
		log.Println("Token not found")
		return &pb.CheckMsgReply{}, nil
	}
	if err != nil {
		log.Println("Failed to get on rdb")
		Health.NowDead()
		return &pb.CheckMsgReply{}, nil
	}
	phone, err := rdb.Get(ctx, "M"+tel+msg).Result()
	if err == redis.Nil {
		log.Println("Phone not found")
		return &pb.CheckMsgReply{}, nil
	}
	if err != nil {
		log.Println("Failed to get on rdb")
		Health.NowDead()
		return &pb.CheckMsgReply{}, nil
	}
	// If phone not found create a new person
	_, err = pdb.QueryContext(ctx, `INSERT INTO people (balance, phone, dlimit, created, last_in)
	VALUES (0, $1, 1, current_timestamp, current_timestamp)
	ON CONFLICT (phone) DO NOTHING;`, phone)
	if err != nil {
		log.Println("Failed in pdb phone check")
		Health.NowDead()
		return &pb.CheckMsgReply{}, nil
	}
	// Get pid
	var pid int64
	err = pdb.QueryRowContext(ctx, `SELECT pid
	FROM people WHERE phone = $1;`, phone).Scan(&pid)
	if err != nil {
		log.Println("Failed to get pid by phone")
		Health.NowDead()
		return &pb.CheckMsgReply{}, nil
	}
	// Set person for session
	err = rdb.Set(ctx, "s"+string(token), string(r.Int64ToByteSlice(pid)), redis.KeepTTL).Err()
	if err != nil {
		log.Println("Failed to set person for session")
		Health.NowDead()
		return &pb.CheckMsgReply{}, nil
	}
	return &pb.CheckMsgReply{Person: pid}, nil
}

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
		log.Fatal("TEL must not be empty")
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
	val, err := rdb.SetNX(ctx, "m"+string(in.GetToken())+tel+msg, "", waitMsgTtl).Result()
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
	return &pb.CheckMsgReply{}, nil
}

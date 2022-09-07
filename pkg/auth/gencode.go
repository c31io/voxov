package auth

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

var maxCode uint64 = 1000000

func byteToCode1(b []byte) uint64 {
	return binary.BigEndian.Uint64(b[:])
}

func byteToCode2(b []byte) uint64 {
	var code uint64
	for _, v := range b {
		code <<= 8
		code |= uint64(v)
	}
	return code
}

func genCode1() (uint64, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}
	code := byteToCode1(b)
	code %= maxCode
	return code, nil
}

func genCode2() (uint64, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}
	code := byteToCode2(b)
	code %= maxCode
	return code, nil
}

func genCode3() (uint64, error) {
	codeBig, err := rand.Int(rand.Reader, big.NewInt(int64(maxCode)))
	if err != nil {
		return 0, err
	}
	code := codeBig.Uint64()
	return code, nil
}

var GenCode func() (uint64, error) = genCode1

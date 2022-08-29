package gate

import (
	"crypto/rand"
	"log"
)

// It takes more than 1.7*10^8 years to randomly guess one
// in 10^12 valid sessions with a rate of 10^18Hz
const tokenBytes = 16

func genToken() []byte {
	b := make([]byte, tokenBytes)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("crypto/rand returned an error")
	}
	return b
}

package gate

import (
	"crypto/rand"
	"errors"
	"log"
)

// It takes more than 1.7*10^8 years to randomly guess one
// in 10^12 valid sessions with a rate of 10^18Hz
const tokenBytes = 16

type Session struct {
	token  []byte
	person []byte
}

var ErrSessionExpired = errors.New("session: expired")

func genToken() []byte {
	b := make([]byte, tokenBytes)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("session: crypto/rand returned an error")
	}
	return b
}

func NewSession(ttl int64) *Session {
	token := genToken()
	return &Session{
		token: token,
	}
}

func (s *Session) AuthAs(person []byte) {
	s.person = person
}

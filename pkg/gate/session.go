package gate

import (
	"crypto/rand"
	"errors"
	"time"
)

// It takes more than 1.7*10^8 years to randomly guess one
// in 10^12 valid sessions with a rate of 10^18Hz
const tokenBytes = 16

type Session struct {
	token         []byte
	expire        int64 // Unix second
	authenticated bool
	person        []byte
}

var ErrSessionExpired = errors.New("session: expired")

func genToken() []byte {
	b := make([]byte, tokenBytes)
	_, err := rand.Read(b)
	if err != nil {
		panic("session: crypto/rand returned an error")
	}
	return b
}

func NewSession(ttl int64) *Session {
	token := genToken()
	expire := time.Now().Unix() + ttl
	return &Session{
		token:  token,
		expire: expire,
	}
}

func (s *Session) IsExpired() bool {
	return s.expire < time.Now().Unix()
}

func (s *Session) Renew(ttl int64) error {
	if s.IsExpired() {
		return ErrSessionExpired
	}
	s.expire = time.Now().Unix() + ttl
	return nil
}

func (s *Session) AuthenticateAs(person []byte) {
	s.authenticated = true
	s.person = person
}

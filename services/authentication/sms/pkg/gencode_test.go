package sms

import (
	"crypto/rand"
	"testing"
)

func TestGenCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := GenCode()
		if err != nil {
			t.Errorf("%d: GenCode() error", i)
		}
		t.Logf("%d: %d", i, got)
		if !(got < maxCode) {
			t.Errorf("%d: GenCode() = %d; want in range [0, maxCode)", i, got)
		}
	}
}

func TestByteToCode(t *testing.T) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		t.Errorf("crypto/rand error")
	}
	if byteToCode1(b) != byteToCode2(b) {
		t.Errorf("ByteToCode functions unmatch")
	}
}

func BenchmarkGenCode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genCode1()
	}
}

func BenchmarkGenCode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genCode2()
	}
}

func BenchmarkGenCode3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genCode3()
	}
}

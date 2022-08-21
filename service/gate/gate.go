package main

import (
	"context"

	"github.com/c31io/voxov/pkg/gate"
)

func main() {
	_ = context.Background()
	_ = gate.Rdb
	_ = new(gate.Session)
}

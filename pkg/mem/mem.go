package mem

import "github.com/bradford-hamilton/byteboy/pkg/cart"

type Memory struct {
	cart *cart.Cart
}

func New() *Memory { return &Memory{} }

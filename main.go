package main

import (
	"fmt"

	"github.com/bradford-hamilton/byteboy/pkg/cpu"
)

func main() {
	cpu := cpu.New()
	fmt.Println(cpu)
}

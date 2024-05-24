package main

import (
	"github.com/bradford-hamilton/byteboy/cmd"
	"github.com/faiface/pixel/pixelgl"
)

// pixelgl needs access to the main thread
func main() {
	pixelgl.Run(runByteBoy)
}

func runByteBoy() {
	cmd.Execute()
}

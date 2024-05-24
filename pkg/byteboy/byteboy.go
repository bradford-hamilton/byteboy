package byteboy

import "github.com/bradford-hamilton/byteboy/pkg/cpu"

const (
	// GameBoy screen 144 x 160
	screenHeight = 144
	screenWidth  = 160

	// FPS
	framesPerSecond = 60
)

// ByteBoy represents the main gameboy structure.
type ByteBoy struct {
	cpu *cpu.CPU

	currentScreen  [screenWidth][screenHeight][3]uint8
	renderedScreen [screenWidth][screenHeight][3]uint8
}

func New() *ByteBoy {
	return &ByteBoy{
		cpu:            cpu.New(),
		currentScreen:  [screenWidth][screenHeight][3]uint8{},
		renderedScreen: [screenWidth][screenHeight][3]uint8{},
	}
}

package byteboy

import (
	"github.com/bradford-hamilton/byteboy/pkg/apu"
	"github.com/bradford-hamilton/byteboy/pkg/cpu"
	"github.com/bradford-hamilton/byteboy/pkg/memory"
	"github.com/bradford-hamilton/byteboy/pkg/ppu"
)

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
	mem *memory.Memory
	apu *apu.APU
	ppu *ppu.PPU

	currentScreen  [screenWidth][screenHeight][3]uint8
	renderedScreen [screenWidth][screenHeight][3]uint8

	paused bool
}

func New() *ByteBoy {
	return &ByteBoy{
		cpu:            cpu.New(),
		mem:            memory.New(),
		apu:            apu.New(),
		ppu:            ppu.New(),
		currentScreen:  [screenWidth][screenHeight][3]uint8{},
		renderedScreen: [screenWidth][screenHeight][3]uint8{},
	}
}

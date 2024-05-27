package byteboy

import (
	"github.com/bradford-hamilton/byteboy/pkg/apu"
	"github.com/bradford-hamilton/byteboy/pkg/cpu"
	"github.com/bradford-hamilton/byteboy/pkg/mem"
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
	mem *mem.MMU
	apu *apu.APU
	ppu *ppu.PPU

	screen         [screenWidth][screenHeight][3]uint8
	screenToRender [screenWidth][screenHeight][3]uint8

	paused bool
}

func New() *ByteBoy {
	return &ByteBoy{
		cpu:            cpu.New(),
		mem:            mem.New(),
		apu:            apu.New(),
		ppu:            ppu.New(),
		screen:         [screenWidth][screenHeight][3]uint8{},
		screenToRender: [screenWidth][screenHeight][3]uint8{},
	}
}

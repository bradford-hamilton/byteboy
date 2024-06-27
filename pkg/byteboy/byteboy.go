package byteboy

import (
	"fmt"

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
	mmu *mem.MMU
	apu *apu.APU
	ppu *ppu.PPU

	opcodeHandlers map[byte]func()

	screen         [screenWidth][screenHeight][3]uint8
	screenToRender [screenWidth][screenHeight][3]uint8

	paused bool
}

func New() *ByteBoy {
	b := &ByteBoy{
		cpu:            cpu.New(),
		mmu:            mem.New(),
		apu:            apu.New(),
		ppu:            ppu.New(),
		screen:         [screenWidth][screenHeight][3]uint8{},
		screenToRender: [screenWidth][screenHeight][3]uint8{},
	}

	b.initializeOpcodeHandlers()

	return b
}

func (b *ByteBoy) initializeOpcodeHandlers() {
	b.opcodeHandlers = map[byte]func(){
		0x00: b.NOOP,
	}
}

func (b *ByteBoy) emulateCycle() {
	opcode := b.fetchOpcode()

	if handler, found := b.opcodeHandlers[opcode]; found {
		handler()
	} else {
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}

// fetchOpcode fetches the next opcode from memory.
func (b *ByteBoy) fetchOpcode() byte {
	opcode := b.mmu.ReadByte(b.cpu.PC)
	b.cpu.PC++

	return opcode
}

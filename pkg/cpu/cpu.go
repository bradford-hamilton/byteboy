package cpu

import "fmt"

// CPU represents a game boy's CPU.
// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html
type CPU struct {
	AF RegisterAF
	BC RegisterPair
	DE RegisterPair
	HL RegisterPair

	SP uint16
	PC uint16
}

func New() *CPU {
	cpu := CPU{
		AF: RegisterAF{A: 0x01, F: 0xB0},
		SP: 0xFFFE,
		PC: 0x0100,
	}
	cpu.BC.SetCombined(0x0013)
	cpu.DE.SetCombined(0x00D8)
	cpu.HL.SetCombined(0x014D)

	return &cpu
}

// RegisterPair represents registers that can be either
// one 16-bit register or two separate 8-bit registers.
type RegisterPair struct {
	High byte // 8-bit high
	Low  byte // 8-bit low
}

// Combined returns the full 16-bit combined register value.
func (rp RegisterPair) Combined() uint16 {
	return uint16(rp.High)<<8 | uint16(rp.Low)
}

func (rp *RegisterPair) SetCombined(value uint16) {
	rp.High = byte(value >> 8)
	rp.Low = byte(value & 0xFF)
}

// RegisterAF combines the A register with flags into a single struct.
type RegisterAF struct {
	A byte // accumulator
	F byte // flags
}

// Combined returns the full 16-bit combined register value for RegisterAF.
func (af RegisterAF) Combined() uint16 {
	return uint16(af.A)<<8 | uint16(af.F)
}

const (
	flagZ = 1 << 7 // Zero flag
	flagN = 1 << 6 // Subtraction flag
	flagH = 1 << 5 // Half Carry flag
	flagC = 1 << 4 // Carry flag
)

// SetZ sets flag Z on/off.
func (af *RegisterAF) SetZ(on bool) {
	if on {
		af.F |= flagZ
	} else {
		af.F &^= flagZ
	}
}

// SetN sets flag N on/off.
func (af *RegisterAF) SetN(on bool) {
	if on {
		af.F |= flagN
	} else {
		af.F &^= flagN
	}
}

// SetH sets flag H on/off.
func (af *RegisterAF) SetH(on bool) {
	if on {
		af.F |= flagH
	} else {
		af.F &^= flagH
	}
}

// SetC sets flag C on/off.
func (af *RegisterAF) SetC(on bool) {
	if on {
		af.F |= flagC
	} else {
		af.F &^= flagC
	}
}

func (af RegisterAF) GetZ() bool { return af.F&flagZ != 0 }
func (af RegisterAF) GetN() bool { return af.F&flagN != 0 }
func (af RegisterAF) GetH() bool { return af.F&flagH != 0 }
func (af RegisterAF) GetC() bool { return af.F&flagC != 0 }

func (cpu *CPU) String() string {
	return fmt.Sprintf(
		"CPU State:\n"+
			"AF: %04X  A: %02X  F: %02X\n"+
			"BC: %04X  B: %02X  C: %02X\n"+
			"DE: %04X  D: %02X  E: %02X\n"+
			"HL: %04X  H: %02X  L: %02X\n"+
			"SP: %04X\n"+
			"PC: %04X",
		cpu.AF.Combined(), cpu.AF.A, cpu.AF.F,
		cpu.BC.Combined(), cpu.BC.High, cpu.BC.Low,
		cpu.DE.Combined(), cpu.DE.High, cpu.DE.Low,
		cpu.HL.Combined(), cpu.HL.High, cpu.HL.Low,
		cpu.SP,
		cpu.PC,
	)
}

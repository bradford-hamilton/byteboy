package cpu

import "fmt"

// CPU represents a Game Boy's CPU.
// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html
type CPU struct {
	A byte
	F byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte

	SP uint16
	PC uint16
}

func New() *CPU {
	cpu := CPU{
		A:  0x01,
		F:  0xB0,
		SP: 0xFFFE,
		PC: 0x0100,
		B:  0x00,
		C:  0x13,
		D:  0x00,
		E:  0xD8,
		H:  0x01,
		L:  0x4D,
	}

	return &cpu
}

// GetAF returns the combined value of registers A and F.
func (cpu *CPU) GetAF() uint16 {
	return get16BitValue(cpu.A, cpu.F)
}

// SetAF sets the combined value of registers A and F.
func (cpu *CPU) SetAF(value uint16) {
	cpu.A, cpu.F = set16BitValue(value)
}

// GetBC returns the combined value of registers B and C.
func (cpu *CPU) GetBC() uint16 {
	return get16BitValue(cpu.B, cpu.C)
}

// SetBC sets the combined value of registers B and C.
func (cpu *CPU) SetBC(value uint16) {
	cpu.B, cpu.C = set16BitValue(value)
}

// GetDE returns the combined value of registers D and E.
func (cpu *CPU) GetDE() uint16 {
	return get16BitValue(cpu.D, cpu.E)
}

// SetDE sets the combined value of registers D and E.
func (cpu *CPU) SetDE(value uint16) {
	cpu.D, cpu.E = set16BitValue(value)
}

// GetHL returns the combined value of registers H and L.
func (cpu *CPU) GetHL() uint16 {
	return get16BitValue(cpu.H, cpu.L)
}

// SetHL sets the combined value of registers H and L.
func (cpu *CPU) SetHL(value uint16) {
	cpu.H, cpu.L = set16BitValue(value)
}

// get16BitValue returns the full 16-bit combined value.
func get16BitValue(high, low byte) uint16 {
	return uint16(high)<<8 | uint16(low)
}

// set16BitValue sets the 16-bit combined register value.
func set16BitValue(value uint16) (byte, byte) {
	high := byte(value >> 8)
	low := byte(value & 0xFF)
	return high, low
}

const (
	flagZ = 1 << 7 // Zero flag
	flagN = 1 << 6 // Subtraction flag
	flagH = 1 << 5 // Half Carry flag
	flagC = 1 << 4 // Carry flag
)

// SetZ sets flag Z on/off.
func (cpu *CPU) SetZ(on bool) {
	if on {
		cpu.F |= flagZ
	} else {
		cpu.F &^= flagZ
	}
}

// SetN sets flag N on/off.
func (cpu *CPU) SetN(on bool) {
	if on {
		cpu.F |= flagN
	} else {
		cpu.F &^= flagN
	}
}

// SetH sets flag H on/off.
func (cpu *CPU) SetH(on bool) {
	if on {
		cpu.F |= flagH
	} else {
		cpu.F &^= flagH
	}
}

// SetC sets flag C on/off.
func (cpu *CPU) SetC(on bool) {
	if on {
		cpu.F |= flagC
	} else {
		cpu.F &^= flagC
	}
}

// GetZ returns the status of flag Z.
func (cpu *CPU) GetZ() bool { return cpu.F&flagZ != 0 }

// GetN returns the status of flag N.
func (cpu *CPU) GetN() bool { return cpu.F&flagN != 0 }

// GetH returns the status of flag H.
func (cpu *CPU) GetH() bool { return cpu.F&flagH != 0 }

// GetC returns the status of flag C.
func (cpu *CPU) GetC() bool { return cpu.F&flagC != 0 }

func (cpu *CPU) String() string {
	return fmt.Sprintf(
		"CPU State:\n"+
			"AF: %04X  A: %02X  F: %02X\n"+
			"BC: %04X  B: %02X  C: %02X\n"+
			"DE: %04X  D: %02X  E: %02X\n"+
			"HL: %04X  H: %02X  L: %02X\n"+
			"SP: %04X\n"+
			"PC: %04X\n"+
			"Flags: Z: %t, N: %t, H: %t, C: %t",
		cpu.GetAF(), cpu.A, cpu.F,
		cpu.GetBC(), cpu.B, cpu.C,
		cpu.GetDE(), cpu.D, cpu.E,
		cpu.GetHL(), cpu.H, cpu.L,
		cpu.SP,
		cpu.PC,
		cpu.GetZ(), cpu.GetN(), cpu.GetH(), cpu.GetC(),
	)
}

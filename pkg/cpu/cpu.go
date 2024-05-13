package cpu

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

const (
	flagZ = 1 << 7 // Zero flag
	flagN = 1 << 6 // Subtraction flag
	flagH = 1 << 5 // Half Carry flag
	flagC = 1 << 4 // Carry flag
)

func (af *RegisterAF) SetZ(on bool) {
	if on {
		af.F |= flagZ
	} else {
		af.F &^= flagZ
	}
}

func (af *RegisterAF) SetN(on bool) {
	if on {
		af.F |= flagN
	} else {
		af.F &^= flagN
	}
}

func (af *RegisterAF) SetH(on bool) {
	if on {
		af.F |= flagH
	} else {
		af.F &^= flagH
	}
}

func (af *RegisterAF) SetC(on bool) {
	if on {
		af.F |= flagC
	} else {
		af.F &^= flagC
	}
}

package cpu

type CPU struct {
	AF RegisterAF
	BC RegisterPair
	DE RegisterPair
	HL RegisterPair

	SP uint16
	PC uint16
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

func (rp *RegisterPair) SetFull(value uint16) {
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

// Set/Clear functionality for a RegisterAF
func (af *RegisterAF) SetZ()   { af.F |= flagZ }
func (af *RegisterAF) ClearZ() { af.F &^= flagZ }
func (af *RegisterAF) SetN()   { af.F |= flagN }
func (af *RegisterAF) ClearN() { af.F &^= flagN }
func (af *RegisterAF) SetH()   { af.F |= flagH }
func (af *RegisterAF) ClearH() { af.F &^= flagH }
func (af *RegisterAF) SetC()   { af.F |= flagC }
func (af *RegisterAF) ClearC() { af.F &^= flagC }

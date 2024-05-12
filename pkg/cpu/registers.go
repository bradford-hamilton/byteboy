package cpu

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

func (rp *RegisterPair) Set(value uint16) {
	rp.High = byte(value >> 8)
	rp.Low = byte(value & 0xFF)
}

const (
	flagZ = 7 // Zero Flag
	flagN = 6 // Subtraction Flag
	flagH = 5 // Half Carry Flag
	flagC = 4 // Carry Flag
)

// AFRegister combines the A register with flags into a single struct.
type AFRegister struct {
	A byte // accumulator
	F byte // flags
}

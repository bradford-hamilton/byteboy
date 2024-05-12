package cpu

type CPU struct {
	AF AFRegister
	BC RegisterPair
	DE RegisterPair
	HL RegisterPair

	SP uint16
	PC uint16
}

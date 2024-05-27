package mem

// MMU represents the GameBoy's memory management unit.
type MMU struct {
	// 32 KB ROM total (00-7FFF)
	// - 16 KiB ROM bank 00 - From cartridge, usually a fixed bank
	// - 16 KiB ROM Bank 01–NN - From cartridge, switchable bank via mapper (if any)
	rom [0x8000]byte

	// 8 KB VRAM (8000-9FFF) - In CGB mode, switchable bank 0/1
	vram [0x2000]byte

	// 8 KB External RAM (A000-BFFF) - From cartridge, switchable bank if any
	eram [0x2000]byte

	// 8 KB Work RAM (C000-DFFF)
	// - 4 KiB Work RAM (WRAM): C000 -> CFFF
	// - 4 KiB Work RAM (WRAM): D000 -> DFFF - In CGB mode, switchable bank 1–7
	wram [0x2000]byte

	// Object Attribute Memory (FE00-FE9F)
	oam [0xA0]byte

	// High RAM (FF80-FFFE)
	hram [0x80]byte

	// Interrupt Enable register (FFFF)
	ie byte

	// TODO:
	// I/O Registers (FF00-FF7F)
	// ioReg [0x80]byte
}

func New() *MMU { return &MMU{} }

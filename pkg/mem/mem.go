package mem

import "fmt"

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

func New() *MMU {
	return &MMU{
		rom:  [32768]byte{},
		vram: [8192]byte{},
		eram: [8192]byte{},
		wram: [8192]byte{},
		oam:  [160]byte{},
		hram: [128]byte{},
		ie:   0,
	}
}

// ReadByte reads a byte from the given address.
func (m *MMU) ReadByte(addr uint16) byte {
	switch {
	case addr <= 0x3FFF:
		return m.rom[addr]
	case addr <= 0x7FFF:
		return m.rom[addr]
	case addr <= 0x9FFF:
		return m.vram[addr-0x8000]
	case addr <= 0xBFFF:
		return m.eram[addr-0xA000]
	case addr <= 0xDFFF:
		return m.wram[addr-0xC000]
	case addr <= 0xFDFF:
		return m.wram[addr-0xE000]
	case addr <= 0xFE9F:
		return m.oam[addr-0xFE00]
	case addr <= 0xFEFF:
		// Not usable
		fmt.Println("TODO")
		return 0xFF
	case addr <= 0xFF7F:
		// TODO: I/O registers
		return 0xFF
	case addr <= 0xFFFE:
		return m.hram[addr-0xFF80]
	case addr == 0xFFFF:
		return m.ie
	default:
		return 0xFF
	}
}

// WriteByte writes a byte to the given address.
func (m *MMU) WriteByte(addr uint16, value byte) {
	switch {
	case addr <= 0x7FFF:
		// ROM is typically not writable, but handle as needed for mappers.
	case addr <= 0x9FFF:
		m.vram[addr-0x8000] = value
	case addr <= 0xBFFF:
		m.eram[addr-0xA000] = value
	case addr <= 0xDFFF:
		m.wram[addr-0xC000] = value
	case addr <= 0xFDFF:
		m.wram[addr-0xE000] = value
	case addr <= 0xFE9F:
		m.oam[addr-0xFE00] = value
	case addr <= 0xFEFF:
		// Not usable
		fmt.Println("TODO")
	case addr <= 0xFF7F:
		// TODO: I/O registers
	case addr <= 0xFFFE:
		m.hram[addr-0xFF80] = value
	case addr == 0xFFFF:
		m.ie = value
	}
}

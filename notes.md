- soc -> CPU + PPU + APU
- 8kb VRAM
- 8kb WRAM
- 6 registers, 16 bit each: AF, BC, DE, HL, SP and PC
- Game Boy SoCs contain a Sharp SM83 CPU core, which is an 8-bit CPU core closely resembling
Zilog Z80 and Intel 8080 but is neither

- Memory Map
The Game Boy has a 16-bit address bus, which is used to address ROM, RAM, and I/O.

Start	End	    Description	                    Notes
0000	3FFF	16 KiB ROM bank 00	            From cartridge, usually a fixed bank
4000	7FFF	16 KiB ROM Bank 01–NN	        From cartridge, switchable bank via mapper (if any)
8000	9FFF	8 KiB Video RAM (VRAM)	        In CGB mode, switchable bank 0/1
A000	BFFF	8 KiB External RAM	            From cartridge, switchable bank if any
C000	CFFF	4 KiB Work RAM (WRAM)	
D000	DFFF	4 KiB Work RAM (WRAM)	        In CGB mode, switchable bank 1–7
E000	FDFF	Echo RAM (mirror of C000–DDFF)	Nintendo says use of this area is prohibited.
FE00	FE9F	Object attribute memory (OAM)	
FEA0	FEFF	Not Usable	                    Nintendo says use of this area is prohibited.
FF00	FF7F	I/O Registers
FF80	FFFE	High RAM (HRAM)
FFFF	FFFF	Interrupt Enable register (IE)

- I/O Ranges
The Game Boy uses the following I/O ranges:

Start	End	    First appeared	Purpose
$FF00		    DMG	            Joypad input
$FF01	$FF02	DMG	            Serial transfer
$FF04	$FF07	DMG	            Timer and divider
$FF10	$FF26	DMG	            Audio
$FF30	$FF3F	DMG	            Wave pattern
$FF40	$FF4B	DMG	            LCD Control, Status, Position, Scrolling, and Palettes
$FF4F		    CGB	            VRAM Bank Select
$FF50		    DMG	            Set to non-zero to disable boot ROM
$FF51	$FF55	CGB	            VRAM DMA
$FF68	$FF6B	CGB	            BG / OBJ Palettes
$FF70		    CGB	            WRAM Bank Select

// Package chip8 implements a CHIP-8 emulator.
package chip8

type State struct {
	Memory [4096]byte

	IndexRegister  uint16
	ProgramCounter uint16
}

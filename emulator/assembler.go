package emulator

// emulator of C256 cpu
type u256 [32]byte


type StatusRegister struct {
	RegsNames uint32
	// Accumulator
	// Overflow
	// Operand1
	// Operand2
	// Loop Counter
	// Memory Pointer (4 of them: in1,in2,out1,out2)

}

	// bitwise 2 operands
	"AND",
	"OR",
	"XOR",

	// bitwise accumulator only
	"NOT",
	"NEG",
	"ASL",
	"ASR",
	"LSR",
	"ROL",
	"ROR",

	// branches (from status register)
	// overflow
	// NAN
	// carry
	// sign
	// zero


type Registers struct {
	basic    [16]u256
	status   u256
	codePtr  uint32
	segments [8]Segment
}


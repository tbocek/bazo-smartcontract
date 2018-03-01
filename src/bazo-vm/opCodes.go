package bazo_vm

const (
	// Constants
	NOP   byte = iota // 0x00
	PUSH  byte = iota // 0x01
	POP   byte = iota // 0x02
	LOAD  byte = iota // 0x03
	STORE byte = iota // 0x04

	// Arithmetic
	ADD byte = iota // 0x05
	SUB byte = iota // 0x06
	MUL byte = iota // 0x07
	DIV byte = iota // 0x08
	MOD byte = iota // 0x09
	AND byte = iota // 0x0a
	EQ  byte = iota // 0x0b
	NEQ byte = iota // 0x0c
	LT  byte = iota // 0x0d
	GT  byte = iota // 0x0e
	LTE byte = iota // 0x0f
	GTE byte = iota // 0x10

	// Flow control
	//TODO

	// Stack
	//TODO

	// Crypto
	//TODO

	// Array
	//TODO

	// Exception
	//TODO
)

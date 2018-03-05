package main

const (
	// Push Operations
	PUSH0 byte = iota // 0x0.
	PUSH1 byte = iota // 0x0.
	PUSH2 byte = iota // 0x0.
	PUSH3 byte = iota // 0x0.

	// Arithmetic
	ADD byte = iota // 0x0.
	SUB byte = iota // 0x0.
	MUL byte = iota // 0x0.
	DIV byte = iota // 0x0.
	MOD byte = iota // 0x0.
	AND byte = iota // 0x0.
	OR  byte = iota // 0x0.
	EQ  byte = iota // 0x0.
	NEQ byte = iota // 0x0.
	LT  byte = iota // 0x0.
	GT  byte = iota // 0x0.
	LTE byte = iota // 0x0.
	GTE byte = iota // 0x0.

	// Flow control
	NOP  byte = iota // 0x0.
	JMP  byte = iota // 0x0.
	JMPT byte = iota // 0x0.
	JMPF byte = iota // 0x0.
	CALL byte = iota // 0x0.
	RET  byte = iota // 0x0.

	// Stack
	POP   byte = iota // 0x0.
	LOAD  byte = iota // 0x0.
	STORE byte = iota // 0x0.

	// Crypto
	SHA3 byte = iota // 0x0.

	// Array
	//TODO

	// Exception
	THROW byte = iota // 0x0.
)

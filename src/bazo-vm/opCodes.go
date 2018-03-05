package main

const (
	// Push Operations
	PUSH0 byte = iota // 0x00 - Empty array of bytes is pushed onto the stack
	PUSH1 byte = iota // 0x01 -
	PUSH2 byte = iota // 0x02
	PUSH3 byte = iota // 0x03
	PUSH4 byte = iota // 0x04

	// Arithmetic
	ADD byte = iota // 0x05 - Add two
	SUB byte = iota // 0x06 -
	MUL byte = iota // 0x07
	DIV byte = iota // 0x08
	MOD byte = iota // 0x09
	AND byte = iota // 0x0a
	OR  byte = iota // 0x0b
	EQ  byte = iota // 0x0c
	NEQ byte = iota // 0x0d
	LT  byte = iota // 0x0e
	GT  byte = iota // 0x0f
	LTE byte = iota // 0x10
	GTE byte = iota // 0x11

	// Flow control
	NOP  byte = iota // 0x12
	JMP  byte = iota // 0x13
	JMPT byte = iota // 0x14
	JMPF byte = iota // 0x15
	CALL byte = iota // 0x16
	RET  byte = iota // 0x17

	// Stack
	POP   byte = iota // 0x18
	LOAD  byte = iota // 0x19
	STORE byte = iota // 0x1a

	// Crypto
	SHA3 byte = iota // 0x1b

	// Array
	//TODO

	// Exception
	THROW byte = iota // 0x1c
)

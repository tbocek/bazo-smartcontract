package bazo_vm

const (
	PUSH = iota
	DUP
	ROLL
	POP
	ADD
	SUB
	MULT
	DIV
	MOD
	NEG
	EQ
	NEQ
	LT
	GT
	LTE
	GTE
	SHIFTL
	SHIFTR
	NOP
	JMP
	JMPIF
	CALL
	CALLEXT
	RET
	SIZE
	STORE
	SSTORE
	LOAD
	SLOAD
	SHA3
	CHECKSIG
	ERRHALT
	HALT
)

const (
	INT = iota
	BYTE
	BYTES
)

type OpCode struct {
	Name     string
	Nargs    int
	ArgTypes []int
	gasPrice uint64
}

var OpCodes = map[int]OpCode{
	PUSH:     {"push", 1, []int{INT}, 1},
	DUP:      {"dup", 0, []int{}, 1},
	ROLL:     {"roll", 1, []int{INT}, 1},
	POP:      {"pop", 0, []int{}, 1},
	ADD:      {"add", 0, []int{}, 1},
	SUB:      {"sub", 0, []int{}, 1},
	MULT:     {"mult", 0, []int{}, 1},
	DIV:      {"div", 0, []int{}, 1},
	MOD:      {"mod", 0, []int{}, 1},
	NEG:      {"neg", 0, []int{}, 1},
	EQ:       {"eq", 0, []int{}, 1},
	NEQ:      {"neq", 0, []int{}, 1},
	LT:       {"lt", 0, []int{}, 1},
	GT:       {"gt", 0, []int{}, 1},
	LTE:      {"lte", 0, []int{}, 1},
	GTE:      {"gte", 0, []int{}, 1},
	SHIFTL:   {"shiftl", 1, []int{BYTE}, 1},
	SHIFTR:   {"shiftl", 1, []int{BYTE}, 1},
	NOP:      {"nop", 0, []int{}, 1},
	JMP:      {"jmp", 1, []int{BYTE}, 1},
	JMPIF:    {"jmpif", 1, []int{BYTE}, 1},
	CALL:     {"call", 2, []int{BYTE, BYTE}, 1},
	CALLEXT:  {"callext", 3, []int{BYTES, BYTES, BYTE}, 1},
	RET:      {"ret", 0, []int{}, 1},
	SIZE:     {"size", 0, []int{}, 1},
	STORE:    {"store", 0, []int{}, 1},
	SSTORE:   {"sstore", 1, []int{INT}, 1},
	LOAD:     {"load", 1, []int{INT}, 1},
	SLOAD:    {"sload", 1, []int{INT}, 1},
	SHA3:     {"sha3", 0, []int{}, 1},
	CHECKSIG: {"checksig", 0, []int{}, 1},
	HALT:     {"halt", 0, []int{}, 0},
	ERRHALT:  {"errhalt", 0, []int{}, 0},
}

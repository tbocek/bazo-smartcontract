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
	NEWMAP
	MAPPUSH
	MAPGETVAL
	MAPREMOVE
	NEWARR
	ARRAPPEND
	//ARRINSERT
	ARRREMOVE
	ARRAT
	SHA3
	CHECKSIG
	ERRHALT
	HALT
)

type OpCode struct {
	name     string
	nargs    int
	gasPrice uint64
}


var OpCodes = map[int]OpCode{
	PUSH:     OpCode{"push", 0, 1},
	DUP:      OpCode{"dup", 0, 1},
	ROLL:     OpCode{"roll", 1, 1},
	POP:      OpCode{"pop", 0, 1},
	ADD:      OpCode{"add", 0, 1},
	SUB:      OpCode{"sub", 0, 1},
	MULT:     OpCode{"mult", 0, 1},
	DIV:      OpCode{"div", 0, 1},
	MOD:      OpCode{"mod", 0, 1},
	NEG:      OpCode{"neg", 0, 1},
	EQ:       OpCode{"eq", 0, 1},
	NEQ:      OpCode{"neq", 0, 1},
	LT:       OpCode{"lt", 0, 1},
	GT:       OpCode{"gt", 0, 1},
	LTE:      OpCode{"lte", 0, 1},
	GTE:      OpCode{"gte", 0, 1},
	SHIFTL:   OpCode{"shiftl", 1, 1},
	SHIFTR:   OpCode{"shiftl", 1, 1},
	NOP:      OpCode{"nop", 0, 1},
	JMP:      OpCode{"jmp", 1, 1},
	JMPIF:    OpCode{"jmpif", 1, 1},
	CALL:     OpCode{"call", 2, 1},
	CALLEXT:  OpCode{"callext", 3, 1},
	RET:      OpCode{"ret", 0, 1},
	SIZE:     OpCode{"size", 0, 1},
	STORE:    OpCode{"store", 0, 1},
	SSTORE:   OpCode{"sstore", 1, 1},
	LOAD:     OpCode{"load", 1, 1},
	SLOAD:    OpCode{"sload", 1, 1},
	NEWMAP:    OpCode{"newmap", 0, 1},
	MAPPUSH:    OpCode{"mappush", 1, 1},
	MAPGETVAL:    OpCode{"mapgetval", 1, 1},
	MAPREMOVE:    OpCode{"mapremove", 1, 1},
	NEWARR:    OpCode{"newarr", 1, 1},
	ARRAPPEND:    OpCode{"arrappend", 0, 1},
	//ARRINSERT:    OpCode{"arrinsert", 0, 1},
	ARRREMOVE:    OpCode{"arrremove", 1, 1},
	ARRAT:    OpCode{"arrat", 1, 1},
	SHA3:     OpCode{"sha3", 0, 1},
	CHECKSIG: OpCode{"checksig", 0, 1},
	HALT:     OpCode{"halt", 0, 0},
	ERRHALT:  OpCode{"errhalt", 0, 0},
}

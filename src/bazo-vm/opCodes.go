package bazo_vm

const (
	PUSH = iota
	PUSHS
	ADD
	SUB
	MULT
	DIV
	MOD
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
	RET
	STORE
	LOAD
	MSTORE
	MLOAD
	SHA3
	PRINT
	HALT
)

type OpCode struct {
	name  string
	nargs int
}

var OpCodes = map[int]OpCode{
	PUSH:   OpCode{"push", 0},
	PUSHS:  OpCode{"pushs", 0}, //Variable nargs, handled direclty in trace function
	ADD:    OpCode{"add", 0},
	SUB:    OpCode{"sub", 0},
	MULT:   OpCode{"mult", 0},
	DIV:    OpCode{"div", 0},
	MOD:    OpCode{"mod", 0},
	EQ:     OpCode{"eq", 0},
	NEQ:    OpCode{"neq", 0},
	LT:     OpCode{"lt", 0},
	GT:     OpCode{"gt", 0},
	LTE:    OpCode{"lte", 0},
	GTE:    OpCode{"gte", 0},
	SHIFTL: OpCode{"shiftl", 1},
	SHIFTR: OpCode{"shiftl", 1},
	NOP:    OpCode{"nop", 0},
	JMP:    OpCode{"jmp", 1},
	JMPIF:  OpCode{"jmpif", 1},
	CALL:   OpCode{"call", 2},
	RET:    OpCode{"ret", 0},
	STORE:  OpCode{"store", 0},
	LOAD:   OpCode{"load", 1},
	MSTORE: OpCode{"mstore", 0},
	MLOAD:  OpCode{"mload", 1},
	SHA3:   OpCode{"sha3", 0},
	PRINT:  OpCode{"print", 0},
	HALT:   OpCode{"halt", 0},
}

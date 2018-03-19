package bazo_vm

const (
	PUSH = iota
	PUSH1
	PUSH2
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
	JMP
	JMPIF
	SHA3
	PRINT
	HALT
)

type OpCode struct {
	name  string
	nargs int
}

var OpCodes = map[int]OpCode{
	PUSH:   OpCode{"push", 2},
	PUSH1:  OpCode{"push1", 1},
	PUSH2:  OpCode{"push2", 2},
	PUSHS:  OpCode{"pushs", 1},
	ADD:    OpCode{"add", 0},
	SUB:    OpCode{"sub", 0},
	MULT:   OpCode{"mult", 0},
	DIV:    OpCode{"div", 0},
	MOD:    OpCode{"mod", 0},
	EQ:     OpCode{"eq", 1},
	NEQ:    OpCode{"neq", 1},
	LT:     OpCode{"lt", 1},
	GT:     OpCode{"gt", 1},
	LTE:    OpCode{"lte", 1},
	GTE:    OpCode{"gte", 1},
	SHIFTL: OpCode{"shiftl", 1},
	SHIFTR: OpCode{"shiftl", 1},
	JMP:    OpCode{"jmp", 1},
	JMPIF:  OpCode{"jmpif", 1},
	SHA3:   OpCode{"sha3", 1},
	PRINT:  OpCode{"print", 0},
	HALT:   OpCode{"halt", 0},
}

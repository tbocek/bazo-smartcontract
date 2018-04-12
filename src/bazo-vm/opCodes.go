package bazo_vm

const (
	PUSH    = iota // 0
	DUP            // 1
	ROLL           // 2
	ADD            // 3
	SUB            // 4
	MULT           // 5
	DIV            // 6
	MOD            // 7
	NEG            // 8
	EQ             // 9
	NEQ            // a
	LT             // b
	GT             // c
	LTE            // d
	GTE            // e
	SHIFTL         // f
	SHIFTR         // 10
	NOP            // 11
	JMP            // 12
	JMPIF          // 13
	CALL           // 14
	CALLEXT        // 15
	RET            // 16
	SIZE           // 17
	STORE          // 18
	LOAD           // 19
	SHA3           // 1a
	ERRHALT        // 1b
	HALT           // 1c
)

type OpCode struct {
	name     string
	nargs    int
	gasPrice int
}

var OpCodes = map[int]OpCode{
	PUSH:    OpCode{"push", 0, 1},
	DUP:     OpCode{"dup", 0, 1},
	ROLL:    OpCode{"roll", 1, 1},
	ADD:     OpCode{"add", 0, 1},
	SUB:     OpCode{"sub", 0, 1},
	MULT:    OpCode{"mult", 0, 1},
	DIV:     OpCode{"div", 0, 1},
	MOD:     OpCode{"mod", 0, 1},
	NEG:     OpCode{"neg", 0, 1},
	EQ:      OpCode{"eq", 0, 1},
	NEQ:     OpCode{"neq", 0, 1},
	LT:      OpCode{"lt", 0, 1},
	GT:      OpCode{"gt", 0, 1},
	LTE:     OpCode{"lte", 0, 1},
	GTE:     OpCode{"gte", 0, 1},
	SHIFTL:  OpCode{"shiftl", 1, 1},
	SHIFTR:  OpCode{"shiftl", 1, 1},
	NOP:     OpCode{"nop", 0, 1},
	JMP:     OpCode{"jmp", 1, 1},
	JMPIF:   OpCode{"jmpif", 1, 1},
	CALL:    OpCode{"call", 2, 1},
	CALLEXT: OpCode{"callext", 3, 1},
	RET:     OpCode{"ret", 0, 1},
	SIZE:    OpCode{"size", 0, 1},
	STORE:   OpCode{"store", 0, 1},
	LOAD:    OpCode{"load", 1, 1},
	SHA3:    OpCode{"sha3", 0, 1},
	HALT:    OpCode{"halt", 0, 0},
	ERRHALT: OpCode{"errhalt", 0, 0},
}

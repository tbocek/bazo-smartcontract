package parser

const (
	OPCODE = iota
	INT
	BYTE
	BYTES
	LABEL
)

type Token struct {
	tokenType int
	value     string
}

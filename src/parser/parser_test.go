package parser

import (
	"fmt"
	"testing"
)

/*func TestProgram(t *testing.T) {
	s := "# asdf asdf asdf\nPUSH 43155\nPUSH 489\nfunctionName:\n"

	fmt.Println(Program(s))
}*/

func TestProgram2(t *testing.T) {
	s := "PUSH 43155\nPUSH 489\nPUSH 4684723121\nPUSH 12\n"

	fmt.Println(Program(s))
}

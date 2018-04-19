package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParser_ProgramAddNums(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/addNums.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)

	if !bytes.Equal(instructionCode, []byte{0, 0, 5, 0, 0, 5, 4, 32}) {
		t.Errorf("After parsing file it should be {0, 0, 5, 0, 0, 5, 4, 32} but is %v", instructionCode)
	}
}

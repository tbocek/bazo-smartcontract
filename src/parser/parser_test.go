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
	fmt.Println(instructionCode)

	if !bytes.Equal(instructionCode, []byte{0, 0, 8, 0, 0, 5, 4, 45}) {
		t.Errorf("Expected tos to be '0, 0, 8, 0, 0, 5, 4, 45' error message but was %v", instructionCode)
	}
}

func TestParser_ProgrammFunctionCall(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/functionCall.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)
	fmt.Println(instructionCode)

	if !bytes.Equal(instructionCode, []byte{0, 1, 217, 228, 0, 0, 5, 21, 12, 2, 45, 27, 0, 27, 1, 4, 23}) {
		t.Errorf("Expected tos to be '0, 1, 217, 228, 0, 0, 5, 21, 12, 2, 45, 27, 0, 27, 1, 4, 23' error message but was %v", instructionCode)
	}
}

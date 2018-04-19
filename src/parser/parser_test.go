package parser

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestProgramAddNums(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/addNums.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)

	if !reflect.DeepEqual(instructionCode, []byte{0, 0, 5, 0, 0, 5, 4, 32}) {
		t.Errorf("After parsing file it should be {0, 0, 5, 0, 0, 5, 4, 32} but is %v", instructionCode)
	}
}

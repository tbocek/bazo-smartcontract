package parser

import (
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
}

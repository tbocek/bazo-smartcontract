package parser

import (
	"bazo-smartcontract/src/bazo-vm"
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Tokenize(sourceCode string) [][]Token {
	var tokenSet [][]Token
	var addressCounter int
	labels := make(map[string]int)

	lines, err := stringToLines(sourceCode)

	if err != nil {
		panic(err)
	}

	for count, line := range lines {

		// Get a string array of every word in line
		words := strings.Fields(line)

		// If case to ignore empty lines
		if len(words) <= 0 {
			continue
		}

		tokenSet = append(tokenSet, []Token{})

		firstWord := words[0]

		if firstWord == "#" {
			continue
		}

		if firstWord[len(firstWord)-1:] == ":" {
			labels[firstWord[:len(firstWord)-1]] = addressCounter
			continue
		}

		for key := range bazo_vm.OpCodes {
			opCode := bazo_vm.OpCodes[key]

			if firstWord == strings.ToUpper(opCode.Name) {
				err := checkIllegalWordsAfterArguments(opCode.Nargs, words)

				if err != nil {
					fmt.Println(err)
				}

				// Handle opCode with no arguments
				tokenSet[count] = append(tokenSet[count], Token{tokenType: OPCODE, value: strings.ToUpper(opCode.Name)})
				addressCounter++

				// Handle opCode with arguments
				for i := 0; i < opCode.Nargs; i++ {
					tokenSet[count] = append(tokenSet[count], Token{tokenType: opCode.ArgTypes[i], value: words[i+1]})
					addressCounter++
				}

			}

		}
	}
	return tokenSet
}

func Parse(sourceCode string) []byte {
	var instructionSet []byte
	tokenSet := Tokenize(sourceCode)

	for lineCount := range tokenSet {
		if lineCount <= 0 {
			continue
		}

		fmt.Println(tokenSet[lineCount])

	}

	return instructionSet
}

// TODO: push
/*
	instructionSet = append(instructionSet, bazo-vm.PUSH)
	val := new(big.Int)
	val.SetString(words[1], 10)

	length := len(val.Bytes()) - 1

	instructionSet = append(instructionSet, byte(length))
	instructionSet = append(instructionSet, val.Bytes()...)

	addressCounter += length + 3
*/

func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func checkIllegalWordsAfterArguments(expectedCount int, words []string) error {
	if len(words) > expectedCount+1 {
		if words[expectedCount+1] != "#" {
			return errors.New("Illegal word in line")
		}
	}
	return nil
}

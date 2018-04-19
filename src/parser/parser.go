package parser

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/bazo-blockchain/bazo-smartcontract/src/vm"
	"github.com/pkg/errors"
)

func Tokenize(sourceCode string) ([][]Token, map[string]int) {
	var tokenSet [][]Token
	var addressCounter int
	var lineCount int
	labels := make(map[string]int)

	lines, err := stringToLines(sourceCode)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {

		// If case to ignore empty lines
		if len(line) <= 0 {
			continue
		}

		// Get a string array of every word in line
		words := strings.Fields(line)

		firstWord := words[0]

		if firstWord == "#" {
			continue
		}

		if firstWord[len(firstWord)-1:] == ":" {
			labels[firstWord[:len(firstWord)-1]] = addressCounter
			continue
		}

		for key := range vm.OpCodes {
			opCode := vm.OpCodes[key]

			if firstWord == strings.ToUpper(opCode.Name) {
				tokenSet = append(tokenSet, []Token{})

				err := checkIllegalWordsAfterArguments(opCode.Nargs, words)

				if err != nil {
					fmt.Println(err)
				}

				// Handle opCode
				tokenSet[lineCount] = append(tokenSet[lineCount], Token{tokenType: OPCODE, value: strings.ToUpper(opCode.Name)})
				addressCounter++

				// Handle arguments
				for i := 0; i < opCode.Nargs; i++ {
					tokenSet[lineCount] = append(tokenSet[lineCount], Token{tokenType: opCode.ArgTypes[i], value: words[i+1]})
					addressCounter++
				}
				lineCount++
				continue
			}

		}
	}
	return tokenSet, labels
}

func Parse(sourceCode string) []byte {
	var instructionSet []byte
	tokenSet, labels := Tokenize(sourceCode)

	fmt.Println(labels)

	for lineCount := range tokenSet {
		if lineCount < 0 {
			continue
		}

		fmt.Println(tokenSet[lineCount])
	}

	return instructionSet
}

// TODO: push
/*
	instructionSet = append(instructionSet, vm.PUSH)
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

package parser

import (
	"bazo-smartcontract/src/vm"
	"bufio"
	"fmt"
	"math/big"
	"strings"

	"github.com/pkg/errors"
)

func Program(sourceCode string) []byte {
	var instructionSet []byte
	var addressCounter int
	labels := make(map[string]int)

	lines, err := stringToLines(sourceCode)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {

		// Get a string array of every word in line
		words := strings.Fields(line)

		// If case to ignore empty lines
		if len(words) > 0 {
			firstWord := words[0]

			switch firstWord {
			case "PUSH":
				err := checkIllegalWordsAfterArguments(1, words)
				if err != nil {
					fmt.Println(err)
				}

				instructionSet = append(instructionSet, vm.PUSH)
				val := new(big.Int)
				val.SetString(words[1], 10)

				length := len(val.Bytes()) - 1

				instructionSet = append(instructionSet, byte(length))
				instructionSet = append(instructionSet, val.Bytes()...)

				addressCounter += length + 3

			case "CALL":
				err := checkIllegalWordsAfterArguments(2, words)
				if err != nil {
					fmt.Println(err)
				}

				instructionSet = append(instructionSet, vm.CALL)

			case "ADD":
				err := checkIllegalWordsAfterArguments(0, words)
				if err != nil {
					fmt.Println(err)
				}

				instructionSet = append(instructionSet, vm.ADD)
				addressCounter++

			case "HALT":
				err := checkIllegalWordsAfterArguments(0, words)
				if err != nil {
					fmt.Println(err)
				}

				instructionSet = append(instructionSet, vm.HALT)
				addressCounter++

			case "#":
				// Do nothing, comments should be ignored

			default:
				if firstWord[len(line)-1:] == ":" {
					labels[firstWord[:len(firstWord)-1]] = addressCounter
				}

				fmt.Println("Invalid first word")
			}
		}
	}
	return instructionSet
}

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

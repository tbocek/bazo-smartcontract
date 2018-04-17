package parser

import (
	"bazo-smartcontract/src/bazo-vm"
	"bufio"
	"fmt"
	"math/big"
	"strings"
)

func Program(sourceCode string) []byte {
	var instructionSet []byte
	var addressCounter int
	labels := make(map[string]int)

	lines, err := StringToLines(sourceCode)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {

		word := firstWord(line)

		switch word {
		case "PUSH":
			instructionSet = append(instructionSet, bazovm.PUSH)
			val := new(big.Int)
			val.SetString(restOfLine(line), 10)

			length := len(val.Bytes()) - 1

			instructionSet = append(instructionSet, byte(length))
			instructionSet = append(instructionSet, val.Bytes()...)

			addressCounter += length

		case "#":
			fmt.Println("This is a comment and to be ignored")

		default:
			if word[len(line)-1:] == ":" {
				labels[word[:len(word)-1]] = addressCounter
			}

			fmt.Println("Invalid first word")
		}
	}

	fmt.Println(labels)
	return instructionSet
}

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// First Word in line is either an opCode, a '#' to signal the beginning of a comment or a label with ':' as ending char
func firstWord(value string) string {
	for i := range value {
		if value[i] == ' ' {
			return value[0:i]
		}
	}
	return value
}

func restOfLine(value string) string {
	for i := range value {
		if value[i] == ' ' {
			return value[i+1:]
		}
	}
	return value
}

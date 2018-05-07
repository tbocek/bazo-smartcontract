package parser

import (
	"bufio"
	"strings"

	"math/big"

	"log"

	"github.com/bazo-blockchain/bazo-smartcontract/src/vm"
	"github.com/pkg/errors"
)

func Parse(sourceCode string) []byte {
	var instructionSet []byte
	tokenSet, labels := Tokenize(sourceCode)

	for lineCount := range tokenSet {
		tokensInLine := tokenSet[lineCount]
		for _, token := range tokensInLine {
			switch token.tokenType {
			case OPCODE:
				opCode, err := getOpCodeIndex(token.value)

				if err != nil {
					log.Fatal(err)
				}

				instructionSet = append(instructionSet, byte(opCode))

			case INT:
				val := new(big.Int)
				val.SetString(token.value, 10)

				if val.String() == "0" {
					instructionSet = append(instructionSet, 0)
					instructionSet = append(instructionSet, 0)
					continue
				}

				length := len(val.Bytes()) - 1

				instructionSet = append(instructionSet, byte(length))
				instructionSet = append(instructionSet, val.Bytes()...)

			case BYTE:
				val := new(big.Int)
				val.SetString(token.value, 10)

				if val.String() == "0" {
					instructionSet = append(instructionSet, 0)
					continue
				}

				instructionSet = append(instructionSet, val.Bytes()...)

			case BYTES:
				val := new(big.Int)
				val.SetString(token.value, 16)

				instructionSet = append(instructionSet, val.Bytes()...)

			case LABEL:
				address := labels[token.value]

				val := new(big.Int)
				val.SetInt64(int64(address))

				instructionSet = append(instructionSet, val.Bytes()[0])
			}
		}
	}
	return instructionSet
}

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
			labels[firstWord[:len(firstWord)-1]] = addressCounter - 2
			continue
		}

		key, err := getOpCodeIndex(firstWord)

		if err != nil {
			log.Fatal(err)
		}

		opCode := vm.OpCodes[key]

		tokenSet = append(tokenSet, []Token{})

		err = checkIllegalWordsAfterArguments(opCode.Nargs, words)

		if err != nil {
			log.Fatal(err, lineCount+1)
		}

		if len(words) <= opCode.Nargs {
			log.Fatal("Missing arguments on line ", lineCount+1)
		}

		// Handle opCode
		tokenSet[lineCount] = append(tokenSet[lineCount], Token{tokenType: OPCODE, value: strings.ToUpper(opCode.Name)})
		addressCounter++

		// Handle arguments if opCode has any
		for i := 0; i < opCode.Nargs; i++ {
			tokenSet[lineCount] = append(tokenSet[lineCount], Token{tokenType: opCode.ArgTypes[i], value: words[i+1]})
			addressCounter++
		}

		// Handle variable int length
		if opCode.Name == "push" {
			val := new(big.Int)
			val.SetString(words[1], 10)

			if val.String() == "0" {
				addressCounter += 2
			} else {
				length := len(val.Bytes())
				addressCounter += length + 1
			}
		}
		lineCount++
	}
	return tokenSet, labels
}

func getOpCodeIndex(s string) (int, error) {
	for key := range vm.OpCodes {
		opCode := vm.OpCodes[key]

		if s == strings.ToUpper(opCode.Name) {
			return key, nil
			continue
		}
	}
	return -1, errors.New("No matching opCode found")
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

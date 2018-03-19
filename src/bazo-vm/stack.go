package bazo_vm

import (
	"errors"
	"log"
)

type byteArray []byte

type Stack struct {
	stack []byteArray
}

func NewStack() Stack {
	return Stack{
		stack: []byteArray{},
	}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(element []byte) {
	s.stack = append(s.stack, element)
}

func (s *Stack) Pop() (element []byte) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("Pop() on empty stack"))
		return []byte{}
	}
}

func (s *Stack) Peek() (element []byte, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return []byte{}, errors.New("Peek() on empty stack!")
	}
}

// Implement String Method to format Stack printout with data formated according datatypes
/*func (s Stack) String() string {
	result := "["
	firstRun := true
	for _, item := range s.stack {
		if firstRun == false {
			result += ", "
		}
		firstRun = false
		result = fmt.Sprint(result, formatData(item.dataType, item.byteArray))
	}
	result += "]"
	return result
}
*/

package bazo_vm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
)

const (
	INT = iota
	FLOAT
	STRING
)

type byteArray []byte

type stackItem struct {
	dataType  byte
	byteArray byteArray
}

type Stack struct {
	stack []stackItem
}

func NewStack() Stack {
	return Stack{
		stack: []stackItem{},
	}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(dataType byte, element byteArray) {
	s.stack = append(s.stack, stackItem{dataType, element})
}

func (s *Stack) PushInt(element int) {
	(*s).Push(INT, IntToByteArray(element))
}

func (s *Stack) PushStr(element string) {
	(*s).Push(STRING, StrToByteArray(element))
}

func (s *Stack) Pop() (element stackItem) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("Pop() on empty stack"))
		return stackItem{}
	}
}

func (s *Stack) PopInt() (element int) {
	if (*s).GetLength() > 0 {
		element = int(binary.LittleEndian.Uint64((*s).stack[s.GetLength()-1].byteArray))
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("PopInt() on empty stack"))
		return -1
	}
}

func (s *Stack) PopStr() (element string) {
	if (*s).GetLength() > 0 {
		element := string((*s).stack[s.GetLength()-1].byteArray[:])
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("PopStr() on empty stack"))
		return ""
	}
}

func (s *Stack) PeekInt() (element int, err error) {
	if (*s).GetLength() > 0 {
		element = int(binary.LittleEndian.Uint64((*s).stack[s.GetLength()-1].byteArray))
		return element, nil
	} else {
		return -1, errors.New("PeekInt() on empty stack!")
	}
}

func (s *Stack) Peek() (element stackItem, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return stackItem{}, errors.New("Peek() on empty stack!")
	}
}

func (s Stack) String() string {
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

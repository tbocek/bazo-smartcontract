package bazo_vm

import "math/big"

type Frame struct {
	variables     map[int]big.Int
	returnAddress int
}

type CallStack struct {
	values []*Frame
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (cs CallStack) GetLength() int {
	return len(cs.values)
}

func (cs *CallStack) Push(element *Frame) {
	cs.values = append(cs.values[:cs.GetLength()], element)
}

func (cs *CallStack) Pop() *Frame {
	element := (*cs).values[cs.GetLength()-1]
	cs.values = cs.values[:cs.GetLength()-1]
	return element
}

func (cs *CallStack) Peek() *Frame {
	return (*cs).values[cs.GetLength()-1]
}

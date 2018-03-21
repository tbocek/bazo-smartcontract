package bazo_vm

type Frame struct {
	variables     map[int][]byte
	returnAddress int
}

type CallStack struct {
	values []Frame
}

func NewCallStack() CallStack {
	return CallStack{
		values: []Frame{},
	}
}

func (cs CallStack) GetLength() int {
	return len(cs.values)
}

func (cs *CallStack) Push(element Frame) {
	cs.values = append(cs.values[:cs.GetLength()], element)
}

func (cs *CallStack) Pop() Frame {
	element := (*cs).values[cs.GetLength()-1]
	cs.values = cs.values[:cs.GetLength()-1]
	return element
}

func (cs *CallStack) Peek() Frame {
	return (*cs).values[cs.GetLength()-1]
}

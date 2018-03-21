package bazo_vm

import (
	"errors"
)

type Memory struct {
	data []byteArray
}

func NewMemory() Memory {
	return Memory{
		data: []byteArray{},
	}
}

func (m Memory) GetLength() int {
	return len(m.data)
}

func (m *Memory) Store(element []byte) {
	m.data = append(m.data, element)
}

func (m *Memory) Load(index int) (element []byte, err error) {
	if (*m).GetLength() >= index {
		return (*m).data[index], nil
	} else {
		return nil, errors.New("Index out of bounds")
	}
}

func (m *Memory) Update(index int, element []byte) {
	m.data[index] = element
}

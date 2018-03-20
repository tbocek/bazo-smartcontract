package bazo_vm

import (
	"fmt"
	"testing"
)

func TestMemory_NewMemory(t *testing.T) {
	m := NewMemory()
	if m.GetLength() != 0 {
		t.Errorf("Expected memory with size 0 but got %v", m.GetLength())
	}
}

func TestMemory_IndexOutOfBounds(t *testing.T) {
	m := NewMemory()
	val, err := m.Load(5)

	if err == nil {
		t.Errorf("Throw error because val was %v", val)
	}
}

func TestMemory_Store(t *testing.T) {
	m := NewMemory()
	m.Store(IntToByteArray(520))

	ba, _ := m.Load(0)
	fmt.Println(ByteArrayToInt(ba))
	fmt.Println(520)

	if ByteArrayToInt(ba) != 520 {
		t.Errorf("Value of loaded variable should be {123, 95}, but was %v", ba)
	}
}

func TestMemory_Update(t *testing.T) {
	m := NewMemory()
	m.Store(IntToByteArray(520))
	m.Store(IntToByteArray(4651))
	m.Store(IntToByteArray(78951))
	m.Store(IntToByteArray(1))

	m.Update(2, IntToByteArray(50))

	ba, _ := m.Load(2)
	if ByteArrayToInt(ba) != 50 {
		t.Errorf("Value of loaded variable should be 50 after updating the value from 4651 to 50, but was %v", ba)
	}
}

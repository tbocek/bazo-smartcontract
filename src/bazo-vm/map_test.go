package bazo_vm

import (
	"testing"
)

func Test_NewMap(t *testing.T) {
	m := NewMap()

	if len(m) != 3 {
		t.Errorf("Expected a Byte Array with size 3 but got %v", len(m))
	}
}

func TestMap_Push(t *testing.T) {
	m := NewMap()
	k := []byte{0x01,}
	v := []byte{0x64, 0x00}
	err := m.Push(k, v)

	if err != nil {
		t.Errorf("%v", err)
	}

	ba := []byte(m)
	if len(ba) != 10 {
		t.Errorf("Expected a Byte Array with size 10 but got %v", len(ba))
	}

	if ba[5] != 0x01 {
		t.Errorf("Unexpected key value or unexpected internal data structure")
	}

	if ba[8] != 0x64 {
		t.Errorf("Unexpected key value or unexpected internal data structure")
	}

}
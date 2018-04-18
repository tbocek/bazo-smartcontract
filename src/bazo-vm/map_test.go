package bazo_vm

import (
	"testing"
	"bytes"
)

func Test_NewMap(t *testing.T) {
	m := NewMap()

	if len(m) != 3 {
		t.Errorf("Expected a Byte Array with size 3 but got %v", len(m))
	}
}

func TestMap_IncerementSize(t *testing.T) {
	m := NewMap()

	s := BaToUI16(m[1:3])
	if s != 0 {
		t.Errorf("Invalid Array Size, Expected 0 but got %v", s)
	}

	m.IncrementSize()
	si := BaToUI16(m[1:3])
	if si != 1 {
		t.Errorf("Invalid Map Size, Expected 1 after increment but got %v", si)
	}
}

func TestMap_DecrementSize(t *testing.T) {
	a := Array([]byte{0x02, 0x02, 0x00})

	s := BaToUI16(a[1:3])
	if s != 2 {
		t.Errorf("Invalid Array Size, Expected 2 but got %v", s)
	}

	a.DecrementSize()
	sd := BaToUI16(a[1:3])
	if sd != 1 {
		t.Errorf("Invalid Array Size, Expected 1 after decrement but got %v", sd)
	}
}

func TestMap_Append(t *testing.T) {
	m := NewMap()
	k := []byte{0x01,}
	v := []byte{0x64, 0x00}
	err := m.Append(k, v)

	if err != nil {
		t.Errorf("%v", err)
	}

	ba := []byte(m)
	if len(ba) != 10 {
		t.Errorf("Expected a Byte Array with size 10 but got %v", len(ba))
	}

	if ba[5] != 0x01 {
		t.Errorf("Unexpected key or unexpected internal data structure")
	}

	if ba[8] != 0x64 {
		t.Errorf("Unexpected key or unexpected internal data structure")
	}
}

func TestMap_GetVal(t *testing.T) {
	m := NewMap()
	m.Append([]byte{0x00,}, []byte{0x00,})
	m.Append([]byte{0x01,}, []byte{0x01, 0x01})
	m.Append([]byte{0x02,0x00}, []byte{0x02, 0x02, 0x02})
	m.Append([]byte{0x03,0x00,0x00}, []byte{0x03, 0x03, 0x03, 0x03, 0x03})

	expected4 := []byte{0x03, 0x03, 0x03, 0x03, 0x03}
	actual4, err4 := m.GetVal([]byte{0x03,0x00,0x00})
	if err4 != nil {
		t.Errorf("%v", err4)
	}
	if bytes.Compare(expected4, actual4) != 0{
		t.Errorf("Unexpected value, Expected '%# x' but was '%# x'", expected4, actual4)
	}

	expected3 := []byte{0x02, 0x02, 0x02,}
	actual3, err3 := m.GetVal([]byte{0x02,0x00})
	if err3 != nil {
		t.Errorf("%v", err3)
	}
	if bytes.Compare(expected3, actual3) != 0{
		t.Errorf("Unexpected value, Expected '%# x' but was '%# x'", expected3, actual3)
	}

}
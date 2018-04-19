package bazo_vm

import (
	"testing"
	"math/big"
	"reflect"
	"bytes"
)


func TestNewArray(t *testing.T) {
	a := NewArray()

	if len(a) != 3 {
		t.Errorf("Expected Byte Array with size 3 but got %v", len(a))
	}
}

func TestArrayIncerementSize(t *testing.T) {
	a := NewArray()

	s := BaToUI16(a[1:3])
	if s != 0 {
		t.Errorf("Invalid Array Size, Expected 0 but got %v", s)
	}

	a.IncrementSize()
	si := BaToUI16(a[1:3])
	if si != 1 {
		t.Errorf("Invalid Array Size, Expected 1 after increment but got %v", si)
	}
}

func TestArrayDecrementSize(t *testing.T) {
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

func TestArray_At(t *testing.T) {
	a := Array([]byte{	0x02,
						0x03, 0x00,

						0x08, 0x00,		0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x04, 0x00,		0x65, 0x00, 0x00, 0x00,
						0x02, 0x00, 	0x65, 0x00,
						})

	expected0 := []byte{0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}
	actual0, err0 := a.At(0)
	if err0 != nil {
		t.Errorf("%v", err0)
	}
	if !reflect.DeepEqual(actual0, expected0) {
		t.Errorf("Invalid element, expected '%# x' after append but got '%# x'", expected0, actual0)
	}

	expected1 := []byte{0x65, 0x00, 0x00, 0x00,}
	actual1, err1 := a.At(1)
	if err1 != nil {
		t.Errorf("%v", err1)
	}
	if !reflect.DeepEqual(actual1, expected1) {
		t.Errorf("Invalid element, expected %v after append but got %v", expected1, actual1)
	}

	expected2 := []byte{0x65, 0x00,}
	actual2, err2 := a.At(2)
	if err2 != nil {
		t.Errorf("%v", err2)
	}
	if !reflect.DeepEqual(actual2, expected2) {
		t.Errorf("Invalid element, expected %v after append but got %v", expected2, actual2)
	}

}

func TestArray_Insert(t *testing.T) {
	a := Array([]byte{	0x02,
		0x03, 0x00,

		0x08, 0x00,		0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x04, 0x00,		0x65, 0x00, 0x00, 0x00,
		0x02, 0x00, 	0x65, 0x00,
	})

	v := big.NewInt(1)
	a.Insert(0, *v)

	expected0 := []byte{0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}
	actual0, err0 := a.At(1)
	if err0 != nil {
		t.Errorf("%v", err0)
	}
	if bytes.Compare(actual0, expected0) != 0 {
		t.Errorf("Invalid element, expected '%# x' after insert at pos 0 but got '%# x'", expected0, actual0)
	}

	expected1 := []byte{0x65, 0x00, 0x00, 0x00,}
	actual1, err1 := a.At(2)
	if err1 != nil {
		t.Errorf("%v", err1)
	}
	if bytes.Compare(actual1, expected1) != 0 {
		t.Errorf("Invalid element, expected %v after insert at pos 0 but got %v", expected1, actual1)
	}
}

func TestArray_Append(t *testing.T) {
	a := NewArray()
	el := big.NewInt(12345678910111213)
	err := a.Append(*el)

	if err != nil {
		t.Errorf("%v", err)
	}

	if a.getSize() != 1 {
		t.Errorf("Invalid Array Size, Expected 1 after append but got %v", a.getSize())
	}
}
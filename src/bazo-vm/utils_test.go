package bazo_vm

import (
	"testing"
	"math/big"
	"fmt"
)

func TestIToBA(t *testing.T) {
	ba := UI16ToBa(0)

	if len(ba) != 2 {
		t.Errorf("Expected Byte Array with size 2 but got %v", len(ba))
	}

	var ui16max uint16 = 65535
	ba2 := UI16ToBa(ui16max)

	if uint16(len(ba2)) != 2 {
		t.Errorf("Expected Byte Array with size 2 but got %v", uint16(len(ba2)))
	}
}

func TestBaToUI16(t *testing.T) {
	ba := []byte{0xFF, 0xFF}
	var ui16max uint16 = 65535

	r := BaToUI16(ba)

	if r != ui16max {
		t.Errorf("Expected result to be 65535 but was %v", r)
	}
}

func TestUI16AndBaConversions(t *testing.T){
	ba := UI16ToBa(15)
	r := BaToUI16(ba)
	if r != 15 {
		t.Errorf("Expected result to be 15 but was %v", r)
	}

	ba2 := UI16ToBa(65535)
	r2 := BaToUI16(ba2)
	if r2 != 65535 {
		t.Errorf("Expected result to be 65535 but was %v", r)
	}
}

func TestIntToByteArrayAndBack(t *testing.T) {
	var start uint64 = 4651321
	ba := IToBA(start)

	end := ByteArrayToInt(ba)
	if start != uint64(end) {
		t.Errorf("Converstion from int to byteArray and back failed, start and end should be equal, are start: %v, end: %v", start, end)
	}
}

func TestStrToByteArrayAndBack(t *testing.T) {
	startStr := "asdf"
	ba := StrToBigInt(startStr)

	endStr := BigIntToString(ba)
	if startStr != endStr {
		t.Errorf("Converstion from str to byteArray and back failed, start and end should be equal, are start: %s, end: %s", startStr, endStr)
	}
}

func TestBigIntMap_Marshal(t *testing.T) {

	key := "T"
	value := "t"

	m := bigIntMap{
		Value: map[string]string{
			key: value,
		},
	}

	result := Marshal(m)
	fmt.Println("Result: ", result)
	expected := big.NewInt(0)

	if expected.Cmp(&result) == 0 {
		t.Errorf("Expected marshalled map to be '%v' but was: '%v'", expected, result)
	}
}

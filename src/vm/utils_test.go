package vm

import (
	"testing"
)

func TestUtils_UInt16ToByteArray(t *testing.T) {
	ba := UInt16ToByteArray(0)

	if len(ba) != 2 {
		t.Errorf("Expected Byte Array with size 2 but got %v", len(ba))
	}

	var ui16max uint16 = 65535
	ba2 := UInt16ToByteArray(ui16max)

	if uint16(len(ba2)) != 2 {
		t.Errorf("Expected Byte Array with size 2 but got %v", uint16(len(ba2)))
	}
}

func TestUtils_ByteArrayToUI16(t *testing.T) {
	ba := []byte{0xFF, 0xFF}
	var ui16max uint16 = 65535

	r := ByteArrayToUI16(ba)

	if r != ui16max {
		t.Errorf("Expected result to be 65535 but was %v", r)
	}
}

func TestUtils_UI16AndByteArrayConversions(t *testing.T) {
	ba := UInt16ToByteArray(15)
	r := ByteArrayToUI16(ba)
	if r != 15 {
		t.Errorf("Expected result to be 15 but was %v", r)
	}

	ba2 := UInt16ToByteArray(65535)
	r2 := ByteArrayToUI16(ba2)
	if r2 != 65535 {
		t.Errorf("Expected result to be 65535 but was %v", r)
	}
}

func TestUtils_IntToByteArrayAndBack(t *testing.T) {
	var start uint64 = 4651321
	ba := UInt64ToByteArray(start)

	end := ByteArrayToInt(ba)
	if start != uint64(end) {
		t.Errorf("Converstion from int to byteArray and back failed, start and end should be equal, are start: %v, end: %v", start, end)
	}
}

func TestUtils_StrToByteArrayAndBack(t *testing.T) {
	startStr := "asdf"
	ba := StrToBigInt(startStr)

	endStr := BigIntToString(ba)
	if startStr != endStr {
		t.Errorf("Converstion from str to byteArray and back failed, start and end should be equal, are start: %s, end: %s", startStr, endStr)
	}
}

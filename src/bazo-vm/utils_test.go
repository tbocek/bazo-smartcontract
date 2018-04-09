package bazo_vm

import (
	"testing"
	"math/big"
	"fmt"
)

/*func TestIntToByteArrayAndBack(t *testing.T) {
	var start int = 4651321
	ba := IntToByteArray(start)

	end := ByteArrayToInt(ba)
	if start != end {
		t.Errorf("Converstion from int to byteArray and back failed, start and end should be equal, are start: %v, end: %v", start, end)
	}
}*/

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

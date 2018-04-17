package bazovm

import (
	"testing"
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

package bazo_vm

import (
	"encoding/binary"
)

func IntToByteArray(element int) []byte {
	ba := make([]byte, 8)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba

	//return []byte(strconv.Itoa(element))
}

func StrToByteArray(element string) []byte {
	return []byte(element)
}

func ByteArrayToInt(element []byte) int {
	return int(binary.LittleEndian.Uint64(element))
}

func ByteArrayToString(element []byte) string {
	return string(element[:])
}

/*func formatData(dataType byte, ba []byte) string {
	switch dataType {
	case INT:
		return fmt.Sprint(ByteArrayToInt(ba))
	case STRING:
		return ByteArrayToString(ba)
	default:
		return string(ba)
	}
}
*/

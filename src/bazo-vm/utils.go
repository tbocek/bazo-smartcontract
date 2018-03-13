package bazo_vm

import (
	"encoding/binary"
	"fmt"
)

func IntToByteArray(element int) byteArray {
	ba := make([]byte, 8)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba
}

func StrToByteArray(element string) byteArray {
	return []byte(element)
}

func ByteArrayToInt(element byteArray) int {
	return int(binary.LittleEndian.Uint64(element))
}

func ByteArrayToString(element byteArray) string {
	return string(element[:])
}

func formatData(dataType byte, ba byteArray) string {
	switch dataType {
	case INT:
		return fmt.Sprint(ByteArrayToInt(ba))
	case STRING:
		return ByteArrayToString(ba)
	default:
		return string(ba)
	}
}

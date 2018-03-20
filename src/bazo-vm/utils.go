package bazo_vm

import (
	"encoding/binary"
)

func IntToByteArray(element int) []byte {
	ba := make([]byte, 64)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba
}

func StrToByteArray(element string) []byte {
	return []byte(element)
}

func ByteArrayToInt(element []byte) int {
	ba := make([]byte, 64-len(element))
	ba = append(element, ba...)
	return int(binary.LittleEndian.Uint64(ba))
}

func ByteArrayToString(element []byte) string {
	return string(element[:])
}

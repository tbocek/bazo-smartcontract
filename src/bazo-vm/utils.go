package bazovm

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
)

func IntToByteArray(element int) []byte {
	ba := make([]byte, 64)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba
}

func StrToBigInt(element string) big.Int {
	var result big.Int
	hexEncoded := hex.EncodeToString([]byte(element))
	result.SetString(hexEncoded, 16)
	return result
}

func ByteArrayToInt(element []byte) int {
	ba := make([]byte, 64-len(element))
	ba = append(element, ba...)
	return int(binary.LittleEndian.Uint64(ba))
}

func BigIntToString(element big.Int) string {
	ba := element.Bytes()
	return string(ba[:])
}

func StrToByteArray(element string) []byte {
	return []byte(element)
}

func ByteArrayToString(element []byte) string {
	return string(element[:])
}

package bazo_vm

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"bytes"
	"encoding/gob"
)

/*func IntToByteArray(element int) []byte {
	ba := make([]byte, 64)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba
}*/

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

type bigIntMap struct {
	Value map[string]string
}

func Marshal(m bigIntMap) big.Int{
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(m)
	if err != nil {
		panic(err)
	}

	var result big.Int
	result.SetBytes(b.Bytes())
	return result
}

func (m *bigIntMap) Unmarshal(input big.Int) {

}
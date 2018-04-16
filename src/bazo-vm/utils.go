package bazo_vm

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"bytes"
	"encoding/gob"
)

func IToBA(element uint64) []byte {
	ba := make([]byte, 8)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	return ba
}

func UI16ToBa(element uint16) []byte {
	ba := make([]byte, 2)
	binary.LittleEndian.PutUint16(ba, uint16(element))
	return ba
}

func BaToUI16(element []byte) uint16{
	return binary.LittleEndian.Uint16(element)
}

func StrToBigInt(element string) big.Int {
	var result big.Int
	hexEncoded := hex.EncodeToString([]byte(element))
	result.SetString(hexEncoded, 16)
	return result
}

func BaToi(element []byte) uint64{
	ba := []byte{}
	r := append(ba, element...)
	return binary.LittleEndian.Uint64(r)
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
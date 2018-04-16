package bazo_vm

import (
	"math/big"
	"errors"
)

type Map []byte
func NewMap() Map {
	return []byte{0x01, 0x00, 0x00,}
}

func (m * Map) ToBigInt() big.Int{
	mp := big.Int{}
	mp.SetBytes(*m)
	return mp
}

func (m *Map) Push(key []byte, value []byte) error {
	sk := len(key)
	sv := len(value)
	if sk > int(UINT16_MAX) || sv > int(UINT16_MAX) {
		return errors.New("key or value size overflows uint16")
	}

	tmp := append(*m, UI16ToBa(uint16(sk))...)
	tmp = append(tmp, key...)
	tmp = append(tmp, UI16ToBa(uint16(sv))...)
	tmp = append(tmp, value...)
	*m = tmp
	return nil
}

func (m *Map) GetVal() {

}

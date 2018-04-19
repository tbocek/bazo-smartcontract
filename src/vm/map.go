package vm

import (
	"bytes"
	"errors"
	"math/big"
)

type Map []byte

func NewMap() Map {
	return []byte{0x01, 0x00, 0x00}
}

func (m *Map) ToBigInt() big.Int {
	mp := big.Int{}
	mp.SetBytes(*m)
	return mp
}

func MapFromBigInt(m big.Int) (Map, error) {
	ba := m.Bytes()
	if ba[0] != 0x01 {
		return Map{}, errors.New("invalid datatype supplied")
	}
	return Map(ba), nil
}

func (m *Map) getSize() uint16 {
	return ByteArrayToUI16((*m)[1:3])
}

func (m *Map) setSize(ba []byte) {
	(*m)[1] = ba[0]
	(*m)[2] = ba[1]
}

func (m *Map) IncrementSize() {
	s := m.getSize()
	s++
	m.setSize(UInt16ToByteArray(s))
}

func (m *Map) DecrementSize() error {
	s := m.getSize()

	if s <= 0 {
		return errors.New("Map size already 0")
	}
	s--
	m.setSize(UInt16ToByteArray(s))
	return nil
}

func (m *Map) Append(key []byte, value []byte) error {
	sk := len(key)
	sv := len(value)
	if sk > int(UINT16_MAX) || sv > int(UINT16_MAX) {
		return errors.New("key or value size overflows uint16")
	}

	tmp := append(*m, UInt16ToByteArray(uint16(sk))...)
	tmp = append(tmp, key...)
	tmp = append(tmp, UInt16ToByteArray(uint16(sv))...)
	tmp = append(tmp, value...)
	*m = tmp
	m.IncrementSize()
	return nil
}

func (m *Map) GetVal(key []byte) ([]byte, error) {
	offset := 3
	l := len(*m)
	for i := offset; i < l; {
		if i > l-3 {
			return []byte{}, errors.New("value sizes are 0")
		}

		ksize := ByteArrayToUI16((*m)[i : i+2])

		valueSizeStart := i + 2 + int(ksize)

		k := (*m)[i+2 : valueSizeStart]
		vsize := ByteArrayToUI16((*m)[valueSizeStart : valueSizeStart+2])
		vSizeEnd := valueSizeStart + 2 + int(vsize)
		v := (*m)[valueSizeStart+2 : vSizeEnd]
		if bytes.Equal(key, k) {
			return v, nil
		}

		if i == vSizeEnd {
			return []byte{}, errors.New("value sizes are 0")
		}
		i = vSizeEnd
	}

	return []byte{}, errors.New("key not found")
}

func (m *Map) Remove(key []byte) error {
	offset := 3
	l := len(*m)
	for i := offset; i < l; {
		if i > l-3 {
			return errors.New("value sizes are 0")
		}

		ksize := ByteArrayToUI16((*m)[i : i+2])

		valueSizeStart := i + 2 + int(ksize)

		k := (*m)[i+2 : valueSizeStart]
		vsize := ByteArrayToUI16((*m)[valueSizeStart : valueSizeStart+2])
		vSizeEnd := valueSizeStart + 2 + int(vsize)
		if bytes.Equal(key, k) {
			tmp := append([]byte{}, (*m)[:i]...)
			*m = append(tmp, (*m)[vSizeEnd:]...)
			m.DecrementSize()
			return nil
		}

		if i == vSizeEnd {
			return errors.New("value sizes are 0")
		}
		i = vSizeEnd
	}

	return errors.New("key not found")
}

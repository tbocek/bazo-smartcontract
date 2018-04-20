package vm

import (
	"errors"
	"math/big"
)

type Array []byte

func NewArray() Array {
	ba := []byte{0x02}
	size := []byte{0x00, 0x00}
	return append(ba, size...)
}

func ArrayFromBigInt(arr big.Int) (Array, error) {
	ba := arr.Bytes()
	if ba[0] != 0x02 {
		return Array{}, errors.New("invalid data type supplied")
	}
	return Array(ba), nil
}

func (a *Array) ToBigInt() big.Int {
	arr := big.Int{}
	arr.SetBytes(*a)
	return arr
}

func (a *Array) getSize() uint16 {
	return ByteArrayToUI16((*a)[1:3])
}

func (a *Array) setSize(ba []byte) {
	(*a)[1] = ba[0]
	(*a)[2] = ba[1]
}

func (a *Array) IncrementSize() {
	s := a.getSize()
	s++
	a.setSize(UInt16ToByteArray(s))
}

func (a *Array) DecrementSize() error {
	s := a.getSize()

	if s <= 0 {
		return errors.New("Array size already 0")
	}
	s--
	a.setSize(UInt16ToByteArray(s))
	return nil
}

func (a *Array) At(index uint16) ([]byte, error) {
	var offset uint16 = 3

	if a.getSize() < index {
		return []byte{}, errors.New("array index out of bounds")
	}

	var i uint16 = 0
	var k uint16 = offset
	for ; k < uint16(len(*a)) && i <= index; i++ {
		s := ByteArrayToUI16((*a)[k : k+2])
		if i == index {
			return (*a)[k+2 : k+2+s], nil
		}
		k += 2 + s
	}

	return []byte{}, errors.New("array internals error")
}

func (a *Array) Insert(index uint16, e big.Int) error {
	var offset uint16 = 3

	if a.getSize() < index {
		return errors.New("array index out of bounds")
	}

	var i uint16 = 0
	var k uint16 = offset
	for ; k < uint16(len(*a)) && i <= index; i++ {
		s := ByteArrayToUI16((*a)[k : k+2])
		if i == index {
			tmp := Array{}
			tmp = append(tmp, (*a)[:k]...)
			tmp.Append(e)
			*a = append(tmp, (*a)[k:]...)
			return nil
		}
		k += 2 + s
	}

	return errors.New("array internals error")
}

func (a *Array) Append(e big.Int) error {
	ba := e.Bytes()
	s := len(ba)

	if s > int(UINT16_MAX) {
		return errors.New("Element Size overflow")
	}

	sb := UInt16ToByteArray(uint16(len(ba)))
	*a = append(*a, append(sb, ba...)...)
	a.IncrementSize()
	return nil
}

func (a *Array) Remove(index uint16) error {
	var offset uint16 = 3

	if a.getSize() < index {
		return errors.New("array index out of bounds")
	}

	var i uint16 = 0
	var k uint16 = offset
	for ; k < uint16(len(*a)) && i <= index; i++ {
		s := ByteArrayToUI16((*a)[k : k+2])
		if i == index {
			tmp := Array{}
			tmp = append(tmp, (*a)[:k]...)
			*a = append(tmp, (*a)[k+2+s:]...)
			return nil
		}
		k += 2 + s
	}

	return errors.New("array internals error")
}

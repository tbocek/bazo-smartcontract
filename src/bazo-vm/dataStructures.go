package bazo_vm

import (
	"math/big"
	"errors"
	"fmt"
)

const UINT16_MAX uint16 = 65535

type Array []byte
func NewArray() Array{
	ba := []byte{0x02,}
	size := []byte{0x00, 0x00,}
	return append(ba, size...)
}

func (a * Array) ToBigInt() big.Int{
	arr := big.Int{}
	arr.SetBytes(*a)
	return arr
}

func (a * Array) getSize() uint16{
	return BaToUI16((*a)[1:3])
}

func (a * Array) setSize(ba []byte) {
	(*a)[1] = ba[0]
	(*a)[2] = ba[1]
}

func (a * Array) IncrementSize(){
	s := a.getSize()
	s++
	a.setSize(UI16ToBa(s))
}

func (a * Array) DecrementSize(){
	s := a.getSize()
	s--
	a.setSize(UI16ToBa(s))
}

func (a * Array) At(index uint16) ([]byte, error) {
	var offset uint16 = 3

	if a.getSize() < index {
		return []byte{}, errors.New("array index out of bounds")
	}

	var i uint16 = 0
	var k uint16 = offset
	for ; k < uint16(len(*a)) && i <= index; i++{
		s := BaToUI16((*a)[k:k+2])
		fmt.Println("Size", s)
		fmt.Println("K: ", k)
		if i == index {
			return (*a)[k+2:k+2+s], nil
		}
		k += 2 + s
	}

	return []byte{}, errors.New("array internals error")

}

func (a * Array) Append(e big.Int) error {
	ba := e.Bytes()
	s := len(ba)

	if s > 65535 {
		return errors.New("Element Size overflow")
	}

	sb := UI16ToBa(uint16(len(ba)))
	*a = append(*a, append(sb, ba...)...)
	a.IncrementSize()
	return nil
}











type FixedElementSizeArray []byte
func NewFixedElementSizeArray(elementSize uint16) FixedElementSizeArray{
	ba := []byte{0x002,}
	size := []byte{0x00, 0x00,}
	ba = append(ba, size...)

	es := UI16ToBa(elementSize)
	return append(ba, es...)
}
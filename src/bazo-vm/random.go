package bazovm

import (
	rand1 "crypto/rand"
	rand2 "math/rand"
	"time"
)

func RandomBytes() []byte {
	byteArray := make([]byte, RandomInt())
	rand1.Read(byteArray)
	return byteArray
}

func RandomInt() int {
	rand2.Seed(time.Now().Unix())
	return rand2.Intn(1000)
}

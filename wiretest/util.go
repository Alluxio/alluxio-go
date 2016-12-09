package wiretest

import (
	"math/rand"
)

// RandomBool generates a random bool.
func RandomBool() bool {
	return rand.Int()%2 == 0
}

// RandomBytes generates a random byte slice.
func RandomBytes() []byte {
	result := make([]byte, rand.Intn(128))
	for i := 0; i < len(result); i++ {
		result[i] = byte(rand.Int() & 0xFF)
	}
	return result
}

// RandomString generates a random string.
func RandomString() string {
	result, length := "", rand.Intn(128)
	for i := 0; i < length; i++ {
		result += string(rand.Intn(96) + 33)
	}
	return result
}

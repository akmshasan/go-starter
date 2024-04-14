package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(rand.Intn(26) + 97)
	}
	return string(b)
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomColor() string {
	colors := []string{"RED", "WHITE", "BLACK", "ORANGE", "GREEN", "PURPLE", "CYAN"}
	return colors[RandInt(0, len(colors)-1)]
}

func RandomName() string {
	return RandString(5)
}

func RandomPrice() int {
	return int(RandInt(1, 10))
}

func RandomQuantity() int {
	return int(RandInt(1, 100))
}

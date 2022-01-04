package utils

import (
	"math/rand"
	"time"
)

func GetRandomTimeRequest() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return (r.Intn(10) + 1)
}

package utils

import "math/rand"

const (
	seed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	return string(b)
}

func RandomInt(n int) int {
	return rand.Intn(n)
}

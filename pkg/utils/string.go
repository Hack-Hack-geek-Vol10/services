package utils

import "math/rand"

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[RandomInt(len(letterBytes))]
	}
	return string(b)
}

func RandomInt(n int) int {
	return rand.Intn(n)
}

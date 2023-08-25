package utils

import "math/rand"

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomURL(size int) string {
	stt := make([]rune, size)

	for i := range stt {
		stt[i] = runes[rand.Intn(len(runes))]
	}

	return string(stt)
}

package utils

import (
	"math/rand"
	"strings"
	"time"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func RandomString(length int) string {
	var result strings.Builder
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result.WriteRune(characters[rand.Intn(len(characters))])
	}

	return result.String()
}

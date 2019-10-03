package aol

import (
	"math/rand"
	"strings"
	"time"
)

type RandPassword struct {
	Value string
}

func NewRandPassword(length int) RandPassword {

	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteRune(chars[rand.Intn(len(chars))])
	}
	RandPassword := RandPassword{
		Value: builder.String(),
	}

	return RandPassword
}

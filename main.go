package main

import (
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

const letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func makeRandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 24)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func writeToClip(word string) {
	clipboard.WriteAll(word)
}

func main() {
	randomString := makeRandomString()
	writeToClip(randomString)
}

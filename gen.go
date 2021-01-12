package main

import "crypto/rand"

var HALF = 32768
var RandomBytes = GenerateRandomBytes(HALF)
var LenRandomBytes = int64(len(RandomBytes))

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

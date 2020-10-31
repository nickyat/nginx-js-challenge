package main

import (
	"math/rand"
	"strconv"
)

// genNewNonce generates new cryptographic nonce
func genNewNonce() string {
	mu.Lock()
	number := strconv.Itoa(rand.Intn(maxNonce)) //nolint: gosec
	mu.Unlock()

	return number
}

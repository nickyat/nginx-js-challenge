package main

//nolint: gosec
import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// getStringHash creates hash as string from input string
func getStringHash(text ...string) string {
	b := []byte(strings.Join(text, ""))
	h := sha1.Sum(b) // nolint: gosec

	return fmt.Sprintf("%x", h)
}

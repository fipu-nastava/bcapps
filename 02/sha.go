package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("Kriptografija je cool!\n"))
	fmt.Printf("%x", sum)
}
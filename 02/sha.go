package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("Kriptografija je zanimljiva!\n"))
	fmt.Printf("%x\n", sum)
}

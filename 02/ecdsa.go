package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Printf("Private key: %x \n", privatekey.D)

	fmt.Printf("Public key: %x %x \n", pubkey.X, pubkey.Y)

	sha := sha256.New()
	message := []byte("Poruka koju valja potpisati")
	signhash := sha.Sum(message)

	fmt.Printf("Hash poruke %x\n", signhash)

	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if serr != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// r.Add(r, big.NewInt(1)) // promjena potpisa

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature: %x\n", signature)

	// Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true
}

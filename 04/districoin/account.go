package districoin

import (
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"bytes"
	"fmt"
)

type Account struct {
	Address Address
	PublicKey ecdsa.PublicKey
	PrivateKey ecdsa.PrivateKey
}

type Address [32]byte

func ComputeAddress(key ecdsa.PublicKey) Address {
	buffer := new(bytes.Buffer)
	buffer.Write(key.X.Bytes())
	buffer.Write(key.Y.Bytes())
	return sha256.Sum256(buffer.Bytes())
}

func NewAccount() (a Account) {

	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	keypair := new(ecdsa.PrivateKey)
	keypair, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	if err != nil {
		panic(err)
	}

	a.PrivateKey = *keypair
	a.PublicKey = keypair.PublicKey
	a.Address = ComputeAddress(a.PublicKey)

	return
}

func (a Account) String() string {
	return fmt.Sprintf("(account:%x)", a.Address[:2])
}

func (a Address) String() string {
	return fmt.Sprintf("(address:%x)", a[:2])
}

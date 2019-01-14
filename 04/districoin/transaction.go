package districoin

import (
	"fmt"
	"crypto/ecdsa"
	"bytes"
	"crypto/rand"
	"github.com/minio/sha256-simd"
	"math/big"
	"encoding/gob"
)

type Transaction struct {
	From Address
	To Address
	Amount int
	Signature []byte
	VerificationKey ecdsa.PublicKey
}

type RS struct {
	R []byte
	S []byte
}

func NewTransaction(from Address, to Address, amount int) (t *Transaction) {
	t = &Transaction{}
	t.From = from
	t.To = to
	t.Amount = amount

	return
}

func (t Transaction) ShortString() string {
	return fmt.Sprintf("(txid: %x)", t.Signature[:4])
}

func (t Transaction) String() string {
	return fmt.Sprintf("(txid: %x, desc: %d coins from %s to %s)", t.Signature[:4], t.Amount, t.From, t.To)
}

func (t *Transaction) Sign(key ecdsa.PrivateKey, keyp ecdsa.PublicKey) {
	/*
	Potpisivanje transakcije privatnim ključem.
	Spremanje javnog ključa kao Verification key.
	 */
	buffer := new(bytes.Buffer)

	buffer.Write(t.From[:])
	buffer.Write(t.To[:])
	buffer.WriteString(fmt.Sprintf("%d", t.Amount))

	signature := sha256.Sum256(buffer.Bytes())
	r, s, err := ecdsa.Sign(rand.Reader, &key, signature[:])
	rs := &RS{r.Bytes(), s.Bytes()}

	if err != nil {
		panic(err)
	}

	buffer = &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)
	enc.Encode(rs)

	t.Signature = buffer.Bytes()
	//fmt.Printf("Signature %x \n", t.Signature)
	t.VerificationKey = keyp
}

func (t* Transaction) Verify() bool {
	/*
	Provjera valjanosti transakcije. Transakcija je valjana ako:
		- adresa pošiljatelja izvedena je od javnog ključa provjeru
	    - mogu verificirati njezin potpis javnim ključem
	 */
	buffer := new(bytes.Buffer)

	buffer.Write(t.From[:])
	buffer.Write(t.To[:])
	buffer.WriteString(fmt.Sprintf("%d", t.Amount))

	signature := sha256.Sum256(buffer.Bytes())

	buffer = &bytes.Buffer{}
	buffer.Write(t.Signature)
	dec := gob.NewDecoder(buffer)
	r := big.NewInt(0)
	s := big.NewInt(0)
	var rs RS
	dec.Decode(&rs)
	r.SetBytes(rs.R)
	s.SetBytes(rs.S)

	isValid := ecdsa.Verify(&t.VerificationKey, signature[:], r, s)

	sender := t.From
	isValid = isValid && sender == ComputeAddress(t.VerificationKey)

	return isValid
}
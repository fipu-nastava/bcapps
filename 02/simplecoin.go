package main

import (
	"fmt"
	"math/rand"
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"time"
)

const AddressLength = 4

type Hash [32]byte
type Address [AddressLength]byte
type Transaction struct {
	From Address
	To Address
	Amount int
}
type Block struct {
	Id int
	Transactions []Transaction
	Hash Hash
	PreviousHash Hash
}

func main() {
	rand.Seed(time.Now().Unix())

	bGenesis := CreateGenesisBlock()

	a1 := Address{1, 1, 1, 1} // neka a1 bude vlasnik blockchaina, može štampati novac
	a2 := GenerateNewAddress()
	a3 := GenerateNewAddress()

	fmt.Printf("%+v \n", a1)
	fmt.Printf("%+v \n", bGenesis)

	// PrintMoneyTransaction(a2, 1000) // ovo mora vratiti "panic"
	t0 := PrintMoneyTransaction(a1, 10000)
	t1 := Transaction{From: a1, To: a2, Amount: 100}
	t2 := Transaction{From: a2, To: a1, Amount: 11}

	fmt.Printf("%+v \n", t0)

	b1 := CreateBlock(bGenesis, t0, t1, t2)
	b1.Print()

	t3 := Transaction{From: a1, To: a2, Amount: 1330}
	t4 := Transaction{From: a2, To: a1, Amount: 10}

	b2 := CreateBlock(b1, t3, t4)
	b2.Print()

	t5 := Transaction{From: a2, To: a3, Amount: 22}
	t6 := Transaction{From: a2, To: a1, Amount: 8}
	t7 := Transaction{From: a3, To: a1, Amount: 123}

	b3 := CreateBlock(b2, t5, t6, t7)
	b3.Print()
}

func CreateGenesisBlock() (b Block) {
	return
}

func CreateBlock(previous Block, txs... Transaction) (b Block) {
	b.Id = previous.Id + 1
	b.PreviousHash = previous.Hash
	b.AddTxs(txs...)
	hash := b.CalculateHash()
	b.Hash = hash

	return
}

func (b* Block) AddTxs(transactions ...Transaction) {
	b.Transactions = append(b.Transactions, transactions...)
}

func (b* Block) CalculateHash() (retval Hash) {

	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(b)
	if err != nil {
		panic(err)
	}
	s := buffer.Bytes()
	retval = sha256.Sum256(s)

	return
}

func (h Hash) String() (a string) {
	a = fmt.Sprintf("%x", h[:4])
	return
}

// Block 1 - Hash dd78... - Details: {Id:1 Transactions:[{From:[0 0 0 0] To:[1 1 1 1] Amount:10000} {From:[1 1 1 1] To:[82 253 252 7] Amount:100} {From:[82 253 252 7] To:[1 1 1 1] Amount:11}] Hash:dd78a0a3 PreviousHash:00000000}
func (b* Block) Print() {
	fmt.Printf("Block %d - Hash %x... - Details: ", b.Id, b.Hash[:2])
	fmt.Printf("%+v\n", *b)
}

func PrintMoneyTransaction(destination Address, amount int) Transaction {
	if destination != [4]byte{1, 1, 1, 1} {
		panic("Fraud detected! Only owner can generate money")
	}
	retval := Transaction{To: destination, Amount: amount}
	return retval
}

func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/rand"
)

// AddressLength - how many bytes will addresses have?
const AddressLength = 4

// Hash is 32 bytes
type Hash [32]byte

// Address is 4 bytes
type Address [AddressLength]byte

// Transaction struct will have: from, to, and amount
type Transaction struct {
	From   Address
	To     Address
	Amount int
}

// Block will have list of transactions
type Block struct {
	ID           int
	Transactions []Transaction
	Hash         Hash
	PreviousHash Hash
	Vanity       string
}

func main() {
	rand.Seed(0) //time.Now().Unix()

	bGenesis := CreateGenesisBlock()

	a1 := Address{1, 1, 1, 1} // neka a1 bude vlasnik blockchaina, može štampati novac
	a2 := GenerateNewAddress()
	a3 := GenerateNewAddress()

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

	fmt.Printf("\n*\n* Final block hash is %x\n*", b3.Hash)
}

// CreateGenesisBlock creates the genesis block
func CreateGenesisBlock() (b Block) {
	b.Vanity = "I am the genesis!"
	return
}

//CreateBlock creates the block given the previous one and transactions
func CreateBlock(previous Block, txs ...Transaction) (b Block) {
	b.ID = previous.ID + 1
	b.PreviousHash = previous.Hash
	b.AddTxs(txs...)
	hash := b.CalculateHash()
	b.Hash = hash

	return
}

// AddTxs will add txs to existing block
func (b *Block) AddTxs(transactions ...Transaction) {
	b.Transactions = append(b.Transactions, transactions...)
}

// CalculateHash will get you the hash of the full block
func (b *Block) CalculateHash() (retval Hash) {

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

// Print the block contents
func (b *Block) Print() {
	fmt.Printf("Block %d - Hash %x... - Details: ", b.ID, b.Hash[:2])
	fmt.Printf("%+v\n", *b)
}

//PrintMoneyTransaction will put some money in the system
func PrintMoneyTransaction(destination Address, amount int) Transaction {
	if destination != [4]byte{1, 1, 1, 1} {
		panic("Fraud detected! Only owner can generate money")
	}
	retval := Transaction{To: destination, Amount: amount}
	return retval
}

//GenerateNewAddress for address creation (random bytes)
func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}

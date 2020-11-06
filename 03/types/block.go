package simplecoin

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/big"
	"time"
)

const miningDifficulty = 20

type Hash [32]byte

type Block struct {
	Id           int
	Transactions []Transaction
	Hash         Hash
	PreviousHash Hash
	Nonce        *big.Int
	Vanity       string
}

var genesis = &Block{Hash: [32]byte{}, Id: 0}

func GetGenesisBlock() *Block {
	return genesis
}

func CreateBlock(previous *Block, txs ...Transaction) (b *Block) {
	b = &Block{}
	b.Id = previous.Id + 1
	b.PreviousHash = previous.Hash
	b.addTxs(txs...)
	hash, err := b.calculateHash()
	b.Hash = hash

	for _, t := range txs {
		if !t.Verify() {
			panic("Transactions not valid")
		}
	}

	if err != nil {
		fmt.Println(err)
	}
	b.Nonce = b.Mine(miningDifficulty)

	return
}

func (b *Block) Mine(difficulty uint) (nonce *big.Int) {
	fmt.Printf("Mining block %d... \n", b.Id)
	defer func(start time.Time) {
		seconds := time.Since(start).Seconds()
		speed := float64(nonce.Uint64()) / seconds / 1000.
		fmt.Printf("   ...took %f seconds at %.2f MH/s \n", time.Since(start).Seconds(), speed)
	}(time.Now())

	var one big.Int
	one.SetInt64(1)

	nonce = new(big.Int)
	for {
		buffer := &bytes.Buffer{}
		buffer.Write(nonce.Bytes())
		buffer.Write(b.Hash[:])
		sum1 := sha256.Sum256(buffer.Bytes())
		sum := sha256.Sum256(sum1[:]) // double sha
		var b = new(big.Int).Lsh(big.NewInt(1), 256-difficulty)
		var s = new(big.Int).SetBytes(sum[:])

		if s.Cmp(b) == -1 {
			return nonce
		}

		nonce.Add(nonce, &one)
	}
}

func (b *Block) addTxs(transactions ...Transaction) {
	b.Transactions = append(b.Transactions, transactions...)
}

func (b *Block) calculateHash() (retval Hash, err error) {
	s, err := b.serialize()
	if err != nil {
		return
	}
	retval = sha256.Sum256(s)
	return
}

func (b *Block) serialize() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	gob.Register(elliptic.P256()) // Used in transaction

	err := enc.Encode(*b)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (b Block) String() (retval string) {
	retval += fmt.Sprintf("(height: %d, id:%x, nonce: %d, txs:", b.Id, b.Hash[:2], b.Nonce)
	for i, t := range b.Transactions {
		retval += t.ShortString()
		if i < len(b.Transactions)-1 {
			retval += ", "
		}
	}
	retval += ")"

	return
}

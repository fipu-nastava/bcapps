package main

import (
	"simplecoin"
)

func main() {
	a1 := simplecoin.NewAccount()
	a2 := simplecoin.NewAccount()

	t1 := simplecoin.NewTransaction(a1.Address, a2.Address, 1000)
	t1.Sign(a1.PrivateKey, a1.PublicKey) // a1 potpisuje

	bGenesis := simplecoin.GetGenesisBlock()
	b1 := simplecoin.CreateBlock(bGenesis, *t1) // creating a new block

	t2 := simplecoin.NewTransaction(a2.Address, a1.Address, 10)
	t2.Sign(a2.PrivateKey, a2.PublicKey)

	b2 := simplecoin.CreateBlock(b1, *t1, *t2)
	_ = b2
}
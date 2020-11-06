package main

import . "github.com/ntankovic/simplecoin/types"

func main() {
	a1 := NewAccount()
	a2 := NewAccount()

	t1 := NewTransaction(a1.Address, a2.Address, 1000)
	t1.Sign(a1.PrivateKey, a1.PublicKey) // a1 potpisuje

	bGenesis := GetGenesisBlock()
	b1 := CreateBlock(bGenesis, *t1) // creating a new block

	t2 := NewTransaction(a2.Address, a1.Address, 10)
	t2.Sign(a2.PrivateKey, a2.PublicKey)

	b2 := CreateBlock(b1, *t1, *t2)
	_ = b2
}

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type BlockHash [32]byte

type Block struct {
	MyHash       BlockHash
	PreviousHash BlockHash
	Data         string
	previous     *Block
}

func (h BlockHash) String() string {
	return fmt.Sprintf("#%x", []byte(h[:2]))
}

func (n Block) String() (retval string) {
	pn := &n
	format := func(pn *Block) string { return fmt.Sprintf("%s(%s)", pn.MyHash, pn.Data) }

	retval += format(pn)

	for pn.previous != nil {
		pn = pn.previous
		retval += fmt.Sprintf("-> %s ", format(pn))
	}
	return
}

func (n *Block) addBlock(s string) (newN Block) {
	newN.Data = s
	newN.previous = n
	newN.PreviousHash = n.MyHash
	newN.MyHash = newN.Hash()

	return
}

func (n *Block) Hash() (retval BlockHash) {
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(*n)
	if err != nil {
		panic(fmt.Sprintf("There was an error creating CalculateHash %+v", err))
	}
	retval = sha256.Sum256(buffer.Bytes())
	return
}

func main() {
	a := Block{Data: "<GENESIS>"}
	b := a.addBlock("a")
	c := b.addBlock("b")
	d := c.addBlock("c")

	a1 := Block{Data: "<GENESIS>"}
	b1 := a1.addBlock("a")
	c1 := b1.addBlock("b")
	d1 := c1.addBlock("с")

	fmt.Printf("Prvi blockchain: %s \n", d)
	fmt.Printf("Drugi blockchain: %s \n", d1)
}

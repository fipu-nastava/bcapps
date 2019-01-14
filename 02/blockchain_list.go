package main

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"fmt"
)

type NodeHash [32]byte

type Node struct {
	MyHash NodeHash
	PreviousHash NodeHash
	Data string
	previous *Node
}

func (h NodeHash) String() string {
	return fmt.Sprintf("#%x", []byte( h[:2] ))
}

func (n Node) String() (retval string) {
	pn := &n
	format := func(pn *Node) string { return fmt.Sprintf("%s(%s)", pn.MyHash, pn.Data) }

	retval += format(pn)

	for pn.previous != nil {
		pn = pn.previous
		retval += fmt.Sprintf("-> %s ", format(pn))
	}
	return
}

func (n *Node) addElement(s string) (newN Node) {
	newN.Data = s
	newN.previous = n
	newN.PreviousHash = n.MyHash
	newN.MyHash = newN.Hash()

	return
}

func (n* Node) Hash() (retval NodeHash) {
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
	a := Node{Data: "<GENESIS>"}
	b := a.addElement("a")
	c := b.addElement("b")
	d := c.addElement("c")

	a1 := Node{Data: "<GENESIS>"}
	b1 := a1.addElement("a")
	c1 := b1.addElement("b")
	d1 := c1.addElement("с")

	fmt.Printf("Prvi blockchain: %s \n", d)
	fmt.Printf("Drugi blockchain: %s \n", d1)
}


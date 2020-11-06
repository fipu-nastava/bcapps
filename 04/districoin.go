package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	. "github.com/ntankovic/districoin/types"
)

func main() {
	rand.Seed(time.Now().Unix())
	port := flag.Int("port", 57000+rand.Intn(1000), "Port where node should listen")
	flag.Parse()

	fmt.Printf("Bootstrap node: %d \n", *port)

	go StartNode("127.0.0.1", *port)

	fmt.Printf("Node started on localhost:%d ... \n", *port)

	a1 := NewAccount()
	a2 := NewAccount()

	t1 := NewTransaction(a1.Address, a2.Address, 1000)
	t1.Sign(a1.PrivateKey, a1.PublicKey) // a1 potpisuje

	bGenesis := GetGenesisBlock()
	b1 := CreateBlock(bGenesis, *t1) // creating a new block
	b2 := CreateBlock(b1, *t1)       // creating a new block

	m1 := &BlockMessage{b1}
	m2 := &BlockMessage{b2}

	if *port > 57000 {
		AddClient("127.0.0.1", 57000)
	}

	for {
		Broadcast(m1)
		time.Sleep(5 * time.Second)
		Broadcast(m2)
		time.Sleep(5 * time.Second)
	}
}

func StartNode(address string, port int) {
	msgChannel := make(chan Message)
	go StartServer(address, port, msgChannel)
	go HandleMessage(msgChannel, port)

}

func HandleMessage(msgChannel chan Message, port int) {
	for {
		m := <-msgChannel
		fmt.Printf("[node:%d] Received message from : %+v \n", port, m)
		bm := m.(*BlockMessage)
		fmt.Println(bm.Block.Id)
	}
}

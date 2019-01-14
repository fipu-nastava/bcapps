package districoin

import (
	"encoding/gob"
	"bytes"
	. "net"
	"fmt"
	"bufio"
)

var clients []*TCPConn

type Message interface {
    Serialize() []byte
}

type BlockMessage struct {
	Block *Block
}

func (b BlockMessage) Serialize() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	enc.Encode(b)

	return buf.Bytes()
}

func AddClient(host string, port int) {
	ip, err := LookupIP(host)
	if err != nil {
		panic(err)
	}

	conn, err := DialTCP("tcp", nil, &TCPAddr{IP:ip[0], Port: port})
	if err != nil {
		panic(err)
	}
	conn.SetKeepAlive(true)
	conn.SetNoDelay(true)
	clients = append(clients, conn)
}

func StartServer(host string, port int, msgChannel chan Message) {
	ip, err := LookupIP(host)
	if err != nil {
		panic(err)
	}

	listener, err := ListenTCP("tcp", &TCPAddr{IP:ip[0], Port: port})
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			for { // there can be many messages
				var m BlockMessage

				r := bufio.NewReader(conn)
				dec := gob.NewDecoder(r)
				err := dec.Decode(&m)

				if err != nil {
					panic(err)
				}

				msgChannel <- Message(&m)
			}
		}()
	}
}

func Broadcast(m Message) {
	b := m.Serialize()

	for _, c := range clients {
		fmt.Printf("Broadcasting message to %s: %x \n", c.RemoteAddr(), b)
		c.Write(b)
	}

	// decoding for test
	buf := &bytes.Buffer{}
	buf.Write(b)
	dec := gob.NewDecoder(buf)
	var decBlock BlockMessage
	dec.Decode(&decBlock)
	//fmt.Printf("(%T) %+v\n ", decBlock, decBlock)
}

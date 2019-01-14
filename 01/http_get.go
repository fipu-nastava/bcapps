package main

import (
	"net"
	"fmt"
	"strings"
	"time"
	"bufio"
)

func timeTrack(start time.Time, name string) {
	fmt.Println(name, "took", time.Since(start).Seconds(), "seconds")
}

func main() {

	host := "example.com"

	var repeats = 100
	var verbose = false

	t1 := parallel(host, repeats, verbose)
	t2 := serial(host, repeats, verbose)

	fmt.Printf("Parallel: %d bytes\nSerial: %d bytes\n", t1, t2)
}

func serial(host string, repeats int, verbose bool) int {

	defer timeTrack(time.Now(), "Serial")
	totalBytes := 0

	for i := 0; i < repeats; i++ {
		totalBytes += get(host, nil, verbose)
	}
	return totalBytes
}

func parallel(host string, repeats int, verbose bool) int {

	defer timeTrack(time.Now(), "Parallel")

	c := make(chan int, repeats)

	for i := 0; i < repeats; i++ {
		go get(host, c, verbose)
	}

	totalBytes := 0
	for i := 0; i < repeats; i++ {
		totalBytes += <-c
	}

	return totalBytes
}

func get(host string, c chan int, verbose bool) int {

	conn, err := net.Dial("tcp", host+":80")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return 0
	}

	httpReq := strings.Join([]string{"GET / HTTP/1.1", fmt.Sprintf("Host: %s", host), "Accept: */*", "", ""}, "\r\n")

	if verbose {
		fmt.Println(httpReq)
	}

	n, err := conn.Write([]byte(httpReq))

	if err != nil {
		fmt.Println(err)
		return 0
	}

	if verbose {
		fmt.Println("Sent", n, "bytes")
	}

	r := bufio.NewReader(conn)
	bytesRead := 0

	for {
		if err != nil {
			fmt.Println(err)
			break
		}
		response, err := r.ReadString('\n')

		bytesRead += len(response)

		if err != nil {
			fmt.Println(err)
		}
		if verbose {
			fmt.Print(response)
		}
		if strings.Contains(response, "</html>") {
			conn.Close()
			break
		}
	}

	if c != nil {
		c <- bytesRead
	}

	return bytesRead
}

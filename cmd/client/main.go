package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "server:5000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
	<-time.After(5 * time.Second)

	w := bufio.NewWriter(conn)
	w.WriteString("test\n")
	w.Flush()
	<-time.After(10 * time.Second)
}

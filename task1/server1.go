package main

import (
	"net"
	"fmt"
	"bufio"
)

func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg,_ := reader.ReadString('\n')
	fmt.Println(msg)
	//sends a message back
	fmt.Fprintln(*conn, "OK")
}

func main(){
	ln,_ := net.Listen("tcp", ":8030")
	//allows read in of multiple messages
	for {
		conn,_ := ln.Accept()
		go handleConnection(&conn)
	}

}

package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

//functionality for a read back from server
func readin(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg,_ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main(){
	//implementation of user input
	stdin := bufio.NewReader(os.Stdin)
	//handles connection
	for {
		fmt.Println("Enter text:")
		text,_ := stdin.ReadString('\n')
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		fmt.Fprintf(conn, text)
		readin(&conn)
	}
}
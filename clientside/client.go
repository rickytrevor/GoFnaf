package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var status string

func main() {
	var ip string
	fmt.Println("give me the ipadress:port")
	fmt.Scan(&ip)
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("wrong adress")
	}

	for {
		go getinput(conn)
		getgamedata(conn)

	}
}
func getinput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(conn, line)

	}
}

func getgamedata(conn net.Conn) {

	reader := bufio.NewReader(conn)
	for {
		in, err := reader.ReadString('\n')
		_ = err
		fmt.Print(in)

	}
}

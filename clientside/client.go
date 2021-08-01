package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var status bool

func main() {
	var ip string
	status = true
	ip = " "
	fmt.Println("give me the address:port")
	fmt.Print("address) ")
	fmt.Scan(&ip)
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("ip/port not found")

	}
	for {
		if status == false {
			fmt.Println("thanks for playing!")
			break
		}
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
		if err != nil {
			break
		}
		if strings.Contains(in, "office") {
			status = false
			break
		}
		fmt.Print(in)

	}
}

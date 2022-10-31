package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

var status bool

func main() {

	var ip string
	status = true
	ip = " "
	var conn net.Conn
	var err error

	fmt.Println(conn, "\033[2J")
	myFigure := figure.NewColorFigure("Fnaf Client", "doom", "green", true)
	myFigure.Print()

	for {
		fmt.Println("\nGive me the address:port ( for example 0.0.0.0:8082 )")
		fmt.Print("\nEnter for 0.0.0.0:8082: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ip = scanner.Text()
		//fmt.Scan(&ip)

		if ip == "" {
			ip = "0.0.0.0:8082"
		}

		conn, err = net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("ip/port not found")
		} else {
			break
		}
	}

	fmt.Println("Successfully connected to " + ip)
	fmt.Println()

	for {
		if status == false {
			fmt.Println("thanks for playing!")
			break
		}

		go getinput(conn)
		getgamedata(conn)

		fmt.Println("Connection closed")
		return

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
		if strings.Contains(in, "has entered the office") {
			status = false
			conn.Close()
			break
		}
		fmt.Print(in)

	}
}

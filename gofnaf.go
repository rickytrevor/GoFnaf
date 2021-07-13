package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var cmd string
var testi string
var ldoor bool = true
var rdoor bool = true
var i int
var battery float32 = 100
var consumption float32 = 0.225

type room struct {
	name         string
	nnearbyrooms int
	nearbyrooms  []int
}
type enemy struct {
	name         string
	intelligence int
	currentroom  int
}

var bonnie = enemy{
	name:         "bonnie",
	intelligence: 14,
	currentroom:  0,
}
var freddy = enemy{
	name:         "freddy",
	intelligence: 15,
	currentroom:  0,
}
var chica = enemy{
	name:         "chica",
	intelligence: 9,
	currentroom:  0,
}
var foxy = enemy{
	name:         "foxy",
	intelligence: 4,
	currentroom:  1,
}

func main() {

	rooms := []room{
		room{"stanza1", 2, []int{1, 2}},
		room{"stanza2", 1, []int{0}},
	}
	// devo RIMUOVERE questa linea sotto
	fmt.Print(rooms[1])
	l, _ := net.Listen("tcp", ":8080")

	defer l.Close()
	fmt.Println("yey il server è partito")

	for {
		conn, err := l.Accept()
		if err != nil {
		}
		go timer()
		go test()
		go getinput(conn)
	}

}
func test() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		fmt.Println("captured:", line)
		if line == "madonna cana" {
			fmt.Println("dio schiofoso maiale")
		}
	}
}
func getinput(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			break
		}
		extline := scanner.Text()

		fmt.Println("captured:", extline)
		switch extline {
		case "close right door":
			fmt.Fprintln(conn, "you've closed the right door")
			rdoor = false
		case "close left door":
			fmt.Fprintln(conn, "you've closed the left door")
			ldoor = false
		case "open left door":
			fmt.Fprintln(conn, "you've opened the left door")
			ldoor = true
		case "open right door":
			fmt.Fprintln(conn, "you've opened the right door")
			rdoor = true
		}
		//move(conn)
		fmt.Println("right door status= ", rdoor)
		fmt.Println("left door status= ", ldoor)
	}
}

func randgen() int {
	min := 0
	max := 20
	fmt.Println(rand.Intn(max-min) + min)
	i = (rand.Intn(max-min) + min)
	return i
}

func move(conn net.Conn) {
	randgen()

	dio := 0
	cane := 5
	moviment := (rand.Intn(cane-dio) + dio)
	switch moviment {
	case 1:
		if bonnie.intelligence <= i {
			fmt.Fprintln(conn, "bonnie has changed room.....")
		} else {
			fmt.Println(conn, "bonnie is still here....")
		}
	case 2:
		if freddy.intelligence <= i {
			fmt.Fprintln(conn, "freddy has changed room.....")
		} else {
			fmt.Println(conn, "freddy is still here....")
		}
	case 3:
		if chica.intelligence <= i {
			fmt.Fprintln(conn, "chica has changed room.....")
		} else {
			fmt.Println(conn, "chica is still here....")
		}

	case 4:
		if foxy.intelligence <= i {
			fmt.Fprintln(conn, "foxy has changed room.....")
		} else {
			fmt.Println(conn, "foxy is still here....")
		}
	}
}
func setconsumption() {

}
func timer() int {
	for {
		start := time.Now()
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
		time.Sleep(2 * time.Second)
		battery = battery - consumption
		fmt.Println(battery)

	}
}

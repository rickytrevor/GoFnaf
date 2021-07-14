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
var ldoorcons float32 = 0
var rdoorcons float32 = 0

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
		//sta sulla jamboard la mappa numerata a modo
		room{"stanza1", 2, []int{1, 2}},
		room{"stanza2", 8, []int{1, 3, 4, 7, 6, 11, 5, 8}},
		room{"stanza3", 1, []int{2}},
		room{"stanza4", 1, []int{2}},
		room{"stanza5", 1, []int{2}},
		room{"stanza6", 2, []int{7, 2}},
		room{"stanza7", 2, []int{2, 6}},
		room{"stanza8", 1, []int{2}},
		// stanza11 è il cesso, mi ero scordato di metterla nel disegnino lol
		room{"stanza11", 1, []int{2}},
		room{"stanza9", 1, []int{8}},
		room{"stanza10", 1, []int{7}},
		room{"critroom1", 1, []int{10}},
		room{"critroom2", 1, []int{9}},
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
			rdoorcons = 0.255
		case "close left door":
			fmt.Fprintln(conn, "you've closed the left door")
			ldoor = false
			ldoorcons = 0.255
		case "open left door":
			fmt.Fprintln(conn, "you've opened the left door")
			ldoor = true
			ldoorcons = 0
		case "open right door":
			fmt.Fprintln(conn, "you've opened the right door")
			rdoor = true
			rdoorcons = 0
		}
		move(conn)
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
	consumption = 0.225
	if ldoor == false {
		consumption = 0.550
	}
	if rdoor == false {
		consumption = 0.750
	}
	if ldoor == false && rdoor == false {
		consumption = 1
	}

}
func timer() int {
	for {
		//setconsumption()
		start := time.Now()
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
		time.Sleep(3 * time.Second)
		battery = battery - (consumption + ldoorcons + rdoorcons)
		fmt.Println(battery)

	}
}
func isattackavalable() {

}

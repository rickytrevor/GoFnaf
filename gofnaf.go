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
var youvelost bool = false

//var enemies []enemy

type room struct {
	id           int
	name         string
	nnearbyrooms int
	nearbyrooms  []int
}
type enemy struct {
	name         string
	intelligence int
	currentroom  int
	allowedrooms []int
	nall         int
}

var bonnie = enemy{
	name:         "bonnie",
	intelligence: 14,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var freddy = enemy{
	name:         "freddy",
	intelligence: 15,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var chica = enemy{
	name:         "chica",
	intelligence: 9,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var foxy = enemy{
	name:         "foxy",
	intelligence: 4,
	currentroom:  1,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}

var enemies []enemy = []enemy{bonnie, chica, freddy, foxy}

var rooms []room = []room{
	room{0, "stage", 1, []int{1}},
	room{1, "mainhall", 8, []int{0, 2, 3, 6, 5, 10, 4, 7}},
	room{2, "spareparts", 1, []int{0}},
	room{3, "piratecove", 1, []int{1}},
	room{4, "kitchen", 1, []int{1}},
	room{5, "laundry", 2, []int{6, 1}},
	room{6, "hall1", 2, []int{1, 5}},
	room{7, "hall2", 2, []int{1, 13}},
	room{8, "restrooms", 1, []int{1}},
	room{9, "stanza9", 2, []int{7, 13}},
	room{10, "stanza10", 1, []int{6}},
	room{11, "critroom1", 1, []int{9}},
	room{12, "critroom2", 1, []int{8}},
	room{13, "office", 2, []int{7, 9}},
}

func main() {

	// devo RIMUOVERE questa linea sotto
	fmt.Print(rooms[1])
	// questa linea starta la connessione non và tolta dio porc
	l, _ := net.Listen("tcp", ":8080")

	defer l.Close()
	fmt.Println("yey il server è partito")

	for {
		conn, err := l.Accept()
		if err != nil {
		}
		go timer(conn)
		go getinput(conn)

	}

}

//questa funzione finchè non implemento il multiplayer è spazio sprecato sul disco
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
		//move(conn)
		fmt.Println("right door status= ", rdoor)
		fmt.Println("left door status= ", ldoor)

	}
}

func randgen() int {
	min := 0
	max := 20
	//fmt.Println(rand.Intn(max-min) + min)
	i = (rand.Intn(max-min) + min)
	return i
}
func newmove() {

	var indice int
	for i := 0; i < 4; i++ {
		e := enemies[i]

		intell := rand.Intn(20)
		if e.intelligence < intell {
			break
		}
		fmt.Println()
		cr := e.currentroom
		r := rooms[cr]
		nrooms := r.nearbyrooms
		nnrooms := r.nnearbyrooms
		for {

			indice = rand.Intn(nnrooms)
			st := nrooms[indice]
			//fmt.Println(st)
			test := false
			for d := 0; d < e.nall; d++ {

				if st == e.allowedrooms[d] {
					test = true
					break
				}
			}
			if test == true {
				break
			}
		}
		newRoom := nrooms[indice]
		enemies[i].currentroom = newRoom
		fmt.Println(enemies[i].name, " si è spostato in ", rooms[newRoom].name)
	}
}
func move(conn net.Conn) {
	randgen()
	newmove()
	/*
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
	*/
}

//probabilmente è inutile sta funzione
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
func timer(conn net.Conn) {
	for {
		//setconsumption()
		start := time.Now()
		t := time.Now()
		elapsed := t.Sub(start)
		_ = elapsed
		// fmt.Println(elapsed)
		time.Sleep(1 * time.Second)
		battery = battery - (0.225 + ldoorcons + rdoorcons)
		move(conn)
		//fmt.Print(youvelost)
		//fmt.Println(battery)
		if youvelost == true {
			break
		}
	}
}
func isattackavalable() {

}

//nonserve
func finalchance() {
	ldoor = true
	rdoor = true
	time.Sleep(2 * time.Second)
	fmt.Println("HAI PERSO")
	youvelost = true
}

func movetorooms(conn net.Conn) {

}

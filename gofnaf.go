package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"math/rand"
	"net"
	"os"
	"strings"
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
var oldroombuf int
var newRoom int
var enemybuf int
var failedattack bool
var timetillend int
var cam string

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
	intelligence: 8,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var freddy = enemy{
	name:         "freddy",
	intelligence: 4,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var chica = enemy{
	//var enemies []enemy
	name:         "chica",
	intelligence: 7,
	currentroom:  0,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var foxy = enemy{
	name:         "foxy",
	intelligence: 7,
	currentroom:  4,
	allowedrooms: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	nall:         14,
}
var enemies []enemy = []enemy{bonnie, chica, freddy, foxy}

//qua definisco come stanno messe le stanze in relazione alle altre
var rooms []room = []room{
	room{0, "stage", 1, []int{1}},
	room{1, "mainhall", 7, []int{2, 3, 4, 5, 6, 7, 9}},
	room{2, "spareparts", 1, []int{1}},
	room{3, "bathrooms", 1, []int{1}},
	room{4, "piratecove", 3, []int{1, 6, 11}},
	room{5, "kitchen", 1, []int{1}},
	room{6, "hall1", 3, []int{1, 7, 8}},
	room{7, "jenitorsroom", 2, []int{6, 8}},
	room{8, "nearoffice1", 1, []int{11}},
	room{9, "hall2", 2, []int{10, 1}},
	room{10, "nearoffice2", 1, []int{12}},
	room{11, "critroom1", 1, []int{13}},
	room{12, "critroom2", 1, []int{13}},
	room{13, "office", 4, []int{6, 9, 1, 4}},
}

func main() {

	fmt.Println(rooms[1].name)

	l, _ := net.Listen("tcp", ":8080")
	defer l.Close()
	fmt.Println("the server has started")
	for {
		conn, err := l.Accept()
		if err != nil {
		}
		go timer(conn)
		go getinput(conn)
	}

}

func test() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		fmt.Println("captured:", line)
		if line == "sus amogus" {
			fmt.Println("test")
		}
	}
}

//questa funzione serve a prendere l'input dal giocatore ed è una merda devo rifarla
func getinput(conn net.Conn) {
	youvelost = false
	failedattack = true
	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "if you don't know the commands type 'help' in the chat")
	for scanner.Scan() {

		if youvelost == true {
			break
		}
		if err := scanner.Err(); err != nil {
			break
		}
		extline := scanner.Text()
		fmt.Println("captured:", extline)
		var camb = " "
		c := strings.Split(extline, "check camera")
		camb = c[len(c)-1]
		cam := strings.TrimSpace(camb)
		fmt.Println(cam)
		if cam != " " {
			for i := 0; i < 4; i++ {
				fmt.Println(enemies[i].name, enemies[i].currentroom)
				enemyroombuf := enemies[i].currentroom
				if rooms[enemyroombuf].name == cam {
					fmt.Fprintln(conn, enemies[i].name+" is in room "+cam)
				}
			}
		}
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
		case "list cameras":
			for i := 0; i < 11; i++ {
				fmt.Fprintln(conn, rooms[i].name)

			}
		case "battery percentage":
			fmt.Fprintln(conn, battery)
		case "help":
			fmt.Fprintln(conn, "open left door ")
			fmt.Fprintln(conn, "close left door ")
			fmt.Fprintln(conn, "open right door ")
			fmt.Fprintln(conn, "close right door ")
			fmt.Fprintln(conn, "list cameras ")
			fmt.Fprintln(conn, "check camera + room name (you can see the room names with list cameras) ")
			fmt.Fprintln(conn, "battery percentage")
		}

		fmt.Println("right door status= ", rdoor)
		fmt.Println("left door status= ", ldoor)

	}
}
func newmove(conn net.Conn) {
	var indice int
	failedattack = true
	for {
		i := rand.Intn(4)
		fmt.Println(i)
		e := enemies[i]
		intell := rand.Intn(20)
		fmt.Println("intelligenza")
		fmt.Println(e.intelligence)
		if e.intelligence <= intell {
			break
		}
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
		enemies[i].currentroom = oldroombuf
		enemies[i].currentroom = newRoom
		//i = enemybuf
		//fmt.Println(newRoom)
		if newRoom == 13 {
			//fmt.Println(oldroombuf)
			if rooms[newRoom].name == "office" && rdoor == false {

				enemies[i].currentroom = 1
				fmt.Println("the attack has failed")
				rooms[newRoom].name = "mainhall"
				failedattack = true

			}
			if rooms[newRoom].name == "office" && ldoor == false {
				fmt.Println("attacco fallito")
				enemies[i].currentroom = 1
				rooms[newRoom].name = "mainhall"
			}
			if rooms[newRoom].name == "office" && rdoor == true {
				fmt.Println("attacco riuscito")
				failedattack = false
			}
			if rooms[newRoom].name == "office" && ldoor == true {
				fmt.Println("attacco riuscito")
				failedattack = false
			}
		}
		if failedattack == false {
			fmt.Println("HAI PERSO MALE")
			youvelost = true
			break
		}
		fmt.Println(enemies[i].name, " si è spostato in ", rooms[newRoom].name)
	}
}

func timer(conn net.Conn) {
	for {

		if youvelost == true {
			break
		}
		//setconsumption()
		start := time.Now()
		t := time.Now()
		elapsed := t.Sub(start)
		_ = elapsed
		// fmt.Println(elapsed)
		time.Sleep(8 * time.Second)
		battery = battery - (0.225 + ldoorcons + rdoorcons)
		//fmt.Fprintln(conn, battery)
		fmt.Println("tick")
		if battery > 0 {
			battery = battery - (0.225 + ldoorcons + rdoorcons)
			//fmt.Fprintln(conn, battery)
		} else {
			ldoor = true
			rdoor = true
		}
		newmove(conn)
		timetillend = timetillend + 1
		if timetillend == 16 {
			fmt.Fprintln(conn, "you've won")

			youvelost = true
		}
	}
}

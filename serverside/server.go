package main

import (
	"bufio"
	"fmt"

	"github.com/common-nighthawk/go-figure"

	_ "fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

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

// ******************************************
// ******************************************
// defining rooms in relation to other rooms*
// ******************************************
// ******************************************
var rooms []room = []room{
	room{0, "stage", 1, []int{1}},
	room{1, "mainhall", 7, []int{2, 3, 4, 5, 6, 7, 9}},
	room{2, "spareparts", 1, []int{1, 7}},
	room{3, "bathrooms", 1, []int{1}},
	room{4, "piratecove", 3, []int{1, 6, 11}},
	room{5, "kitchen", 1, []int{1}},
	room{6, "hall1", 5, []int{1, 4, 7, 8, 13}},
	room{7, "jenitorsroom", 2, []int{6, 8}},
	room{8, "nearoffice1", 1, []int{11}},
	room{9, "hall2", 2, []int{10, 1}},
	room{10, "nearoffice2", 3, []int{6, 12, 13}},
	room{11, "critroom1", 1, []int{13}},
	room{12, "critroom2", 1, []int{13}},
	room{13, "office", 6, []int{6, 9, 1, 4, 6, 10}},
}

func main() {
	number_of_connections := 0
	//startint the server
	port := "8080"
	l, _ := net.Listen("tcp", ":"+port)
	defer l.Close()
	myFigure := figure.NewColorFigure("Five nights", "shadow", "green", true)
	myFigure.Print()
	myFigure1 := figure.NewColorFigure("               at Freddy's", "sblood", "red", true)
	myFigure1.Print()
	myFigure2 := figure.NewFigure("                                 SERVER", "", true)
	myFigure2.Print()

	fmt.Println("Server has started on port " + port)
	for {
		conn, err := l.Accept()
		if err != nil {
		}
		number_of_connections++
		fmt.Println(fmt.Sprintf("current connections: %d", number_of_connections))
		go getinput(conn)
	}

}

// ******************************
// ******************************
// *defining the enemies stats***
// ******************************
// ******************************
func getinput(conn net.Conn) {
	var bonnie = enemy{
		name:         "bonnie",
		intelligence: 12,
		currentroom:  0,
		allowedrooms: []int{0, 1, 2, 6, 7, 8, 11, 13},
		nall:         8,
	}
	var freddy = enemy{
		name:         "freddy",
		intelligence: 10,
		currentroom:  0,
		allowedrooms: []int{0, 1, 9, 10, 13},
		nall:         5,
	}
	var chica = enemy{
		name:         "chica",
		intelligence: 9,
		currentroom:  0,
		allowedrooms: []int{0, 1, 5, 9, 10, 12, 13},
		nall:         7,
	}
	var foxy = enemy{
		name:         "foxy",
		intelligence: 2,
		currentroom:  4,
		allowedrooms: []int{4, 6, 10, 13},
		nall:         4,
	}
	var enemies []enemy = []enemy{bonnie, chica, freddy, foxy}

	var ldoor bool = true
	var rdoor bool = true
	var battery float32 = 100
	var ldoorcons float32 = 0
	var rdoorcons float32 = 0
	var youvelost bool = false
	var oldroombuf int
	var failedattack bool
	var timetillend int

	//*******************************************************
	//*******************************************************
	//*declaring the nested function for the moviment logic**
	//*******************************************************
	//*******************************************************

	newmove := func(conn net.Conn) {
		var indice int
		failedattack = true
		for {
			i := rand.Intn(4)
			fmt.Println(i)
			e := enemies[i]
			intell := rand.Intn(20)
			fmt.Println("smartness")
			fmt.Println(e.intelligence)
			fmt.Println("minimum intelligence level to execute the attack")
			fmt.Println(intell)
			if e.intelligence < intell {
				break
			}
			cr := e.currentroom
			r := rooms[cr]
			nrooms := r.nearbyrooms
			nnrooms := r.nnearbyrooms
			for {

				indice = rand.Intn(nnrooms)
				st := nrooms[indice]
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

			if newRoom == 13 {
				if rooms[newRoom].name == "office" && rdoor == false {

					enemies[i].currentroom = 1
					fmt.Println("the attack has failed")
					rooms[newRoom].name = "mainhall"
					failedattack = true

				}
				if rooms[newRoom].name == "office" && ldoor == false {
					fmt.Println("the attack has failed")
					enemies[i].currentroom = 1
					rooms[newRoom].name = "mainhall"
				}
				if rooms[newRoom].name == "office" && rdoor == true {
					fmt.Println("the attack has succeeded")
					failedattack = false
				}
				if rooms[newRoom].name == "office" && ldoor == true {
					fmt.Println("the attack has succeeded")
					failedattack = false
				}
			}
			if failedattack == false {
				fmt.Fprintln(conn, enemies[i].name, "has entered in the office.....")
				youvelost = true

				break
			}
			fmt.Println("")
			fmt.Println("")
			fmt.Println(enemies[i].name, " has moved in ", rooms[newRoom].name)
			fmt.Println("")
			fmt.Println("")
		}
	}

	title := func(conn net.Conn, t string) {
		fig := figure.NewColorFigure(t, "doom", "green", true)
		fmt.Fprintln(conn, fig.String())
	}

	prompt := func(conn net.Conn) {

		fmt.Fprint(conn, "\n")
		if ldoor {
			fmt.Fprint(conn, "L DOOR OPENED")
		} else {
			fmt.Fprint(conn, "L DOOR CLOSED")
		}

		fmt.Fprint(conn, " | ")
		fmt.Fprint(conn, "battery ", battery, "%")
		fmt.Fprint(conn, " | ")

		if rdoor {
			fmt.Fprint(conn, "R DOOR OPENED")
		} else {
			fmt.Fprint(conn, "R DOOR CLOSED")
		}
		fmt.Fprintln(conn, "\n")
	}

	//*************************************************************
	//*************************************************************
	//*defining the logic behind every game tick and the battery***
	//*************************************************************
	//*************************************************************

	timer := func(conn net.Conn) {
		for {

			if youvelost == true {
				break
			}
			start := time.Now()
			t := time.Now()
			elapsed := t.Sub(start)
			_ = elapsed
			time.Sleep(30 * time.Second)
			fmt.Println("tick")
			if battery > 0 {
				battery = battery - (2 + ldoorcons + rdoorcons)
			} else {
				ldoor = true
				rdoor = true
			}
			newmove(conn)
			timetillend = timetillend + 1
			if timetillend == 40 {
				fmt.Fprintln(conn, "you've won")

				youvelost = true
			}
		}
	}
	//********************************
	//********************************
	//*starting the timer() function**
	//********************************
	//********************************
	go timer(conn)
	youvelost = false
	failedattack = true
	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "Type 'h' for help")
	fmt.Fprintln(conn, "Type Enter to refresh\n")

	fmt.Fprintln(conn, "Good luck!")
	fmt.Fprintln(conn, " ")

	//***********************************************************
	//***********************************************************
	//*defining the logic for the scan of every command in game**
	//***********************************************************
	//***********************************************************

	prompt(conn)

	for scanner.Scan() {

		if youvelost == true {
			conn.Close()
			break
		}
		if err := scanner.Err(); err != nil {
			conn.Close()
			break
		}
		extline := scanner.Text()
		fmt.Println("captured:", extline)

		var camb = " "
		c := strings.Split(extline, "cc")
		camb = c[len(c)-1]
		cam := strings.TrimSpace(camb)
		fmt.Println(cam)
		if cam != " " {
			for i := 0; i < 4; i++ {
				fmt.Println(enemies[i].name, enemies[i].currentroom)
				enemyroombuf := enemies[i].currentroom
				if rooms[enemyroombuf].name == cam {
					fmt.Fprintln(conn, "*******************")
					fmt.Fprintln(conn, enemies[i].name+" is in room "+cam)
					fmt.Fprintln(conn, "*******************")
				}
			}
			fmt.Fprintln(conn, " ")
		}
		switch extline {
		case "cr", "close right door":
			fmt.Fprintln(conn, "you've closed the right door")
			rdoor = false
			rdoorcons = 0.3
		case "cl", "close left door":
			fmt.Fprintln(conn, "you've closed the left door")
			ldoor = false
			ldoorcons = 0.3
		case "ol", "open left door":
			fmt.Fprintln(conn, "you've opened the left door")
			ldoor = true
			ldoorcons = 0
		case "or", "open right door":
			fmt.Fprintln(conn, "you've opened the right door")
			rdoor = true
			rdoorcons = 0
		case "lc", "list cameras":
			title(conn, "List cameras")

			for i := 0; i < 11; i++ {
				fmt.Fprintln(conn, rooms[i].name)

			}
		case "ds", "doors status":
			fmt.Fprintln(conn, "right door status = ", rdoor)
			fmt.Fprintln(conn, "left door status = ", ldoor)

		case "bp", "battery percentage":
			fmt.Fprintln(conn, battery)
		case "lost":
			fmt.Println(youvelost)
		case "m", "map":
			title(conn, "Map")
			fmt.Fprintln(conn, "                                      ┌─────────────────┐")
			fmt.Fprintln(conn, "               BACK                   │      SHOW       │")
			fmt.Fprintln(conn, "                ┌──┐                  │                 │                       BATH")
			fmt.Fprintln(conn, "                │  │ ┌────────────────┴─────────────────┴─────────────────┐  ┌───┐ ┌──┐")
			fmt.Fprintln(conn, "                │  ├─┤                                                    │  │   ├─┤  │")
			fmt.Fprintln(conn, "                │  ├─┤                                                    │  │   ├─┤  │")
			fmt.Fprintln(conn, "                └──┘ │                                                    ├──┤   │ └──┘")
			fmt.Fprintln(conn, "                     │                    DINING AREA                     ├──┤   │")
			fmt.Fprintln(conn, "                ┌──┐ │                                                    │  │   │ ┌──┐")
			fmt.Fprintln(conn, "                │  ├─┤                                                    │  │   ├─┤  │")
			fmt.Fprintln(conn, "                │  ├─┤                                                    │  │   ├─┤  │")
			fmt.Fprintln(conn, "                └──┘ └──────────┬─┬───────────────────────┬─┬────────┬─┬──┘  └───┘ └──┘")
			fmt.Fprintln(conn, "             PIRATE             │ │                       │ │        │ │")
			fmt.Fprintln(conn, "                COVE           ┌┴─┴┐                     ┌┴─┴┐     ┌─┴─┴───────────┐")
			fmt.Fprintln(conn, "                               │   │                     │   │     │    KITCHEN    │")
			fmt.Fprintln(conn, "                               │   │                     │   │     │               │")
			fmt.Fprintln(conn, "                        SUPPLY │   │                     │   │     └───────────────┘")
			fmt.Fprintln(conn, "                         ┌───┐ │   │                     │   │")
			fmt.Fprintln(conn, "                         │   │ │   │ ┌─────────────────┐ │   │")
			fmt.Fprintln(conn, "                         │   ├─┤   │ │                 │ │   │")
			fmt.Fprintln(conn, "                         │   ├─┤   │ │                 │ │   │")
			fmt.Fprintln(conn, "                         │   │ │   ├─┤     OFFICE      ├─┤   │")
			fmt.Fprintln(conn, "                         │   │ │   ├─┤                 ├─┤   │")
			fmt.Fprintln(conn, "                         │   │ │   │ │                 │ │   │")
			fmt.Fprintln(conn, "                         └───┘ └───┘ └─────────────────┘ └───┘")
			fmt.Fprintln(conn, "")

		case "h", "help":
			for i := 0; i < 4; i++ {

			}

			title(conn, "Help")

			fmt.Fprintln(conn, "m: map\n")
			fmt.Fprintln(conn, "lc: list cameras")
			fmt.Fprintln(conn, "cc room: check camera + room name (you can see the room names with list cameras)\n")

			fmt.Fprintln(conn, "ol: open left door")
			fmt.Fprintln(conn, "cl: close left door ")
			fmt.Fprintln(conn, "or: open right door ")
			fmt.Fprintln(conn, "cr: close right door\n")

			fmt.Fprintln(conn, "bp: battery percentage")
			fmt.Println(conn, "ds: doors status")

		}

		prompt(conn)

		fmt.Println("rd: right door status= ", rdoor)
		fmt.Println("ld: left door status= ", ldoor)

	}

}

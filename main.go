package main

import (
	"bufio"
	"fmt"
	"gameoflife/lifeform"
	"image"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
)

var world = lifeform.Newworld(200, 100)
var cycle = 0
var c = ui.NewCanvas()
var step time.Duration = 50
var interval = step * time.Millisecond

//Prints out the array - '.' is the lifeform is alive and blank means the lifeform is dead.
func frame(i [][]lifeform.Lifeform) {
	text := ""
	text += "\n\n\n\n"
	for y := range i {
		for x := range i[y] {
			if i[y][x].Alive == 0 {
				text += " "
			} else {
				text += "."
			}
			if i[y][x].Next == 0 {
				i[y][x].Alive = 0
			} else {
				i[y][x].Alive = 1
			}
		}
		text += "\n"
	}
	fmt.Println(text + "\r")
	fmt.Printf("Cycle: %d \n", cycle)
}

//Prints out the array using termui as a GUI
func guiframe(i [][]lifeform.Lifeform) {
	for y := range i {
		for x := range i[y] {
			if i[y][x].Alive == 0 {
				c.SetPoint(image.Pt(y, x), 0)
			} else {
				c.SetPoint(image.Pt(y, x), 2)
			}
			if i[y][x].Next == 0 {
				i[y][x].Alive = 0
			} else {
				i[y][x].Alive = 1
			}
		}
	}
	ui.Render(c)
	fmt.Printf("Cycle: %d\r", cycle)
}

//Generates a random number
func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

//Generates a random number
func oneOrzero() int {
	z := rand.Intn(2)
	if z < 1 {
		z = 0
	} else {
		z = 1
	}
	return z
}

//Seeds the map with new alive lifeforms randomly
func seed(w *[][]lifeform.Lifeform) {
	world := *w
	for i := 0; i < len(world[0])*200; i++ {
		y := randomNumber(1, len(world)-5)
		x := randomNumber(1, len(world[0])-5)
		world[y][x].Alive = oneOrzero()
		world[y][x].Next = oneOrzero()
	}
	*w = world
}

//Resets the map
func reset(world *[][]lifeform.Lifeform) {
	i := *world
	for y := range i {
		for x := range i[y] {
			i[y][x].Alive = 0
			i[y][x].Next = 0
		}
	}
	*world = i
}

//Main entry
func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	l := []int{2, 3}
	d := []int{3}
Start:
	world = lifeform.Newworld(200, 100)
	game := true
	fmt.Println("Game of Life - press 'n' for new generated world, 'q' to quit, 'g' for GUI low res version, 'r' to change game rules - or press enter to run 1 cycle")
	for game == true {
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "r":
			l = []int{}
			d = []int{}
			fmt.Println("Type in numbers from 0 to 8 for rule B with commas i.e '0,1,2,3': ")
			Scanner.Scan()
			rresult1 := Scanner.Text()
			ruleB := strings.Split(rresult1, ",")
			for indexx := range ruleB {
				integer, _ := strconv.Atoi(ruleB[indexx])
				d = append(d, integer)
			}
			fmt.Println("Now type in numbers from 0 to 8 for rule S with commas i.e '0,3,6,9': ")
			Scanner.Scan()
			rresult2 := Scanner.Text()
			ruleS := strings.Split(rresult2, ",")
			for indexxx := range ruleS {
				integer2, _ := strconv.Atoi(ruleS[indexxx])
				l = append(l, integer2)
			}
			fmt.Println("Rules adjusted!")
			goto Start
		case "n":
			rand.Seed(time.Now().UTC().UnixNano())
			reset(&world)
			seed(&world)
			cycle = 0
		case "q":
			game = false
		case "g":
			fmt.Println("Entering GUI version - press 'q' at any time to return to menu, or 'n' to refresh GUI version")
			Scanner.Scan()
			rand.Seed(time.Now().UTC().UnixNano())
			world = lifeform.Newworld(200, 600)
			seed(&world)
			c.SetRect(0, 0, 200, 600)
			if err := ui.Init(); err != nil {
				log.Fatalf("failed to initialize termui: %v", err)
			}
			uiEvents := ui.PollEvents()
			ticker := time.NewTicker(interval).C
			for {
				select {
				case e := <-uiEvents:
					switch e.ID { // event string/identifier
					case "q", "<C-c>": // press 'q' or 'C-c' to quit
						ui.Close()
						goto Start
					case "n":
						cycle = 0
						reset(&world)
						seed(&world)
					}
				// use Go's built-in tickers for updating and drawing data
				case <-ticker:
					cycle++
					lifeform.Adjust(&world, l, d)
					guiframe(world)
				}
			}
		default:
			lifeform.Adjust(&world, l, d)
			frame(world)
			cycle++
		}
	}
}

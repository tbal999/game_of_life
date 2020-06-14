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
	"github.com/gizak/termui/v3/widgets"
)

var (
	world                  = lifeform.Newworld(200, 100)
	cycle                  = 0
	c                      = ui.NewCanvas()
	p                      = widgets.NewParagraph()
	step     time.Duration = 50
	interval               = step * time.Millisecond
	l                      = []int{2, 3}
	d                      = []int{3}
)

//Prints out the array - '.' is the lifeform is alive and blank means the lifeform is dead.
func frame(i [][]lifeform.Lifeform) {
	text := ""
	text += "\n\n\n\n"
	for y := range i {
		for x := range i[y] {
			if i[y][x].Alive == 0 {
				text += " "
			} else {
				if i[y][x].Still < 4 {
					text += "."
				} else if i[y][x].Still < 30 {
					text += ">"
				} else {
					text += "H"
				}
			}
			if i[y][x].Next == 0 {
				i[y][x].Alive = 0
				if i[y][x].Still > 0 {
					i[y][x].Still--
				}
			} else {
				i[y][x].Alive = 1
				if i[y][x].Still <= 31 {
					i[y][x].Still++
				}
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
				if i[y][x].Still < 4 {
					c.SetPoint(image.Pt(y, x), 7)
				} else if i[y][x].Still < 30 {
					c.SetPoint(image.Pt(y, x), 1)
				} else {
					c.SetPoint(image.Pt(y, x), 2)
				}
			}
			if i[y][x].Next == 0 {
				i[y][x].Alive = 0
				if i[y][x].Still > 0 {
					i[y][x].Still--
				}
			} else {
				i[y][x].Alive = 1
				if i[y][x].Still <= 31 {
					i[y][x].Still++
				}
			}
		}
	}
	out := fmt.Sprintf("Cycle: %d\r", cycle)
	rulez := fmt.Sprintf("B: %d and D: %d", d, l)
	p.Text = out + "\n" + "Rules: \n" + rulez
	ui.Render(c)
	ui.Render(p)
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
Start:
	world = lifeform.Newworld(200, 100)
	game := true
	fmt.Println("Game of Life Cellular Automata - Go")
	fmt.Println("'n' for new generated world")
	fmt.Println("'q' to quit")
	fmt.Println("'g' for GUI low res version")
	fmt.Println("'i' for a list of well known rules & further information")
	fmt.Println("'r' to change game rules")
	fmt.Println("'enter' to run 1 cycle on std out console")
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
		case "i":
			fmt.Println("The B rule is how many life-forms need to be alive and adjacent to any one lifeform for it to come back alive.")
			fmt.Println("The S rule is how many life-forms need to be alive and adjacent to any one lifeform for it to survive.")
			fmt.Println("In order of B rule / S rule:")
			fmt.Println("3/2,3 - Game of Life (default)")
			fmt.Println("3/4,5,6,7,8 - Coral")
			fmt.Println("3,6/2,3 - High Life")
			fmt.Println("5/3,4,5 - Long Life")
			fmt.Println("3,6,7/2,4,5- Move")
			fmt.Println("There are many more online!")
		case "g":
			fmt.Println("Entering GUI version. Menu:")
			fmt.Println("when in GUI mode - press 'q' at any time to return to menu")
			fmt.Println("or press 'n' at any time to refresh GUI version")
			fmt.Println("Press 'enter' to start GUI version...")
			Scanner.Scan()
			rand.Seed(time.Now().UTC().UnixNano())
			world = lifeform.Newworld(200, 600)
			seed(&world)
			p.SetRect(95, 1, 120, 10)
			c.SetRect(1, 1, 90, 50)
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

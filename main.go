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
	world   = lifeform.Newworld(200, 100)
	cycle   = 0
	c       = ui.NewCanvas()
	p       = widgets.NewParagraph()
	l       = []int{2, 3}
	d       = []int{3}
	rate    = 30000
	counter = float64(30)
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
	speed := fmt.Sprintf("%f", counter)
	p.Text = out + "\n" + "Rules: \n" + rulez + "\n" + "Speed: " + speed + " ms"
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

//Seeds the map with new alive lifeforms randomly
func seed(w *[][]lifeform.Lifeform, rate int) {
	world := *w
	for i := 0; i < rate; i++ {
		y := randomNumber(1, len(world)-5)
		x := randomNumber(1, len(world[0])-5)
		if world[y][x].Alive == 0 {
			world[y][x].Alive = 1
		}
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
	world = lifeform.Newworld(200, 100)
	rand.Seed(time.Now().UTC().UnixNano())
	seed(&world, rate)
Start:
	game := true
	fmt.Println(
		`
_________        .__  .__        .__                     
\_   ___ \  ____ |  | |  |  __ __|  | _____ _______      
/    \  \/_/ __ \|  | |  | |  |  \  | \__  \\_  __ \     
\     \___\  ___/|  |_|  |_|  |  /  |__/ __ \|  | \/     
 \______  /\___  >____/____/____/|____(____  /__|        
        \/     \/                          \/            
   _____          __                         __          
  /  _  \  __ ___/  |_  ____   _____ _____ _/  |______   
 /  /_\  \|  |  \   __\/  _ \ /     \\__  \\   __\__  \  
/    |    \  |  /|  | (  <_> )  Y Y  \/ __ \|  |  / __ \_
\____|__  /____/ |__|  \____/|__|_|  (____  /__| (____  /
        \/                         \/     \/          \/ 

		`)
	fmt.Println("'n' for new generated world")
	fmt.Println("'q' to quit")
	fmt.Println("'g' for GUI low res version")
	fmt.Println("'i' for a list of well known rules & further information")
	fmt.Println("'r' to change game rules")
	fmt.Println("'s' to change number multiplier (30000 is standard)")
	fmt.Println("'enter' to run 1 cycle on std out console")
	for game == true {
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "s":
			fmt.Println("Type in new multiplier: ")
			Scanner.Scan()
			rresult0 := Scanner.Text()
			integer0, _ := strconv.Atoi(rresult0)
			rate = integer0
			fmt.Println("Multiplier changed! ")
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
			seed(&world, rate)
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
			fmt.Println("or press 'w'/'s' to slow down / speed the rate of change (in ms)")
			fmt.Println("press 'n' at any time to refresh GUI version")
			fmt.Println("Press 'enter' to start GUI version...")
			Scanner.Scan()
			rand.Seed(time.Now().UTC().UnixNano())
			world = lifeform.Newworld(200, 600)
			seed(&world, rate)
			p.SetRect(95, 1, 120, 10)
			c.SetRect(1, 1, 90, 50)
			if err := ui.Init(); err != nil {
				log.Fatalf("failed to initialize termui: %v", err)
			}
			uiEvents := ui.PollEvents()
			for {
				select {
				case e := <-uiEvents:
					switch e.ID { // event string/identifier
					case "q", "<C-c>": // press 'q' or 'C-c' to quit
						ui.Close()
						goto Start
					case "<MouseLeft>":
						payload := e.Payload.(ui.Mouse)
						x, y := payload.X, payload.Y
						world[y][x].Alive = 1
						world[y][x].Next = 1
						p.Text = fmt.Sprintf("%d, %d", x, y)
					case "n":
						rand.Seed(time.Now().UTC().UnixNano())
						reset(&world)
						seed(&world, rate)
						cycle = 0
					case "s":
						if counter > 0 {
							counter -= 10
						}
					case "w":
						if counter < 1000 {
							counter += 10
						}
					}
				// use Go's built-in tickers for updating and drawing data
				case <-time.After(time.Duration(counter) * time.Millisecond):
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

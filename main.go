package main

import (
	"bufio"
	"fmt"
	"gameoflife/lifeform"
	"math/rand"
	"os"
	"time"
)

//Initialise world - a 30x80 array of lifeforms
var world = lifeform.Newworld(30, 80)

//Prints out the array - X is the lifeform is alive and blank if the lifeform is dead.
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

//Adjusts the state of the lifeform's next cycle depending on whether they're alive or not.
func state(a, b, c, d, e, f, g, h, i lifeform.Lifeform) int {
	total := b.Alive + c.Alive + d.Alive + e.Alive + f.Alive + g.Alive + h.Alive + i.Alive
	switch a.Alive {
	case 1:
		switch total {
		case 2, 3:
			return 1
		case 1, 4, 5, 6, 7, 8:
			return 0
		}
	case 0:
		switch total {
		case 3:
			return 1
		}
	}
	return 0
}

//Seeds the map with new alive lifeforms randomly
func seed(w *[][]lifeform.Lifeform) {
	world := *w
	for i := 0; i < 1200; i++ {
		y := randomNumber(5, 25)
		x := randomNumber(5, 70)
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
	game := true
	fmt.Println("Game of Life - press 'n' for new spawn, 'q' to quit or anything else to run cycle")
	for game == true {
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "n":
			rand.Seed(time.Now().UTC().UnixNano())
			reset(&world)
			seed(&world)
		case "q":
			game = false
		default:
			lifeform.Adjust(&world)
			frame(world)
		}
	}
}

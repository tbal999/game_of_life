package main

import (
	"bufio"
	"fmt"
	"gameoflife/lifeform"
	"math/rand"
	"os"
	"time"
)

//Initialise world - a 30x30 array of lifeforms
var world = lifeform.Newworld(30, 30)

//Prints out the array - X is the lifeform is alive and blank if the lifeform is dead.
func frame(i [][]lifeform.Lifeform) {
	text := ""
	for y := range i {
		for x := range i[y] {
			if i[y][x].Alive == 0 {
				text += " "
			} else {
				text += "X"
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

//Checks to see if the item in array is one away from the edge to prevent indexing errors.
func check(x, y int, z [][]lifeform.Lifeform) bool {
	xaxis := len(z[0]) - 1
	yaxis := len(z) - 1
	if x != 0 && y != 0 {
		if x < xaxis && y < yaxis {
			return true
		}
	}
	return false
}

//Checks each lifeform state
func adjust(world *[][]lifeform.Lifeform) {
	i := *world
	for y := range i {
		for x := range i[y] {
			if check(x, y, i) == true {
				i[y][x].Next = state(i[y][x],
					i[y-1][x-1],
					i[y-1][x],
					i[y-1][x+1],
					i[y+1][x-1],
					i[y+1][x],
					i[y+1][x+1],
					i[y][x-1],
					i[y][x+1])
			}
		}
	}
	*world = i
}

//Generates a random number
func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

//Adjusts the state of the lifeform's next cycle depending on whether they're alive or not.
func state(a, b, c, d, e, f, g, h, i lifeform.Lifeform) int {
	switch a.Alive {
	case 1:
		if b.Alive+c.Alive+d.Alive+e.Alive+f.Alive+g.Alive+h.Alive+i.Alive < 2 {
			return 0
		}
		if b.Alive+c.Alive+d.Alive+e.Alive+f.Alive+g.Alive+h.Alive+i.Alive > 3 {
			return 0
		}
		if b.Alive+c.Alive+d.Alive+e.Alive+f.Alive+g.Alive+h.Alive+i.Alive == 2 {
			return 1
		}
		if b.Alive+c.Alive+d.Alive+e.Alive+f.Alive+g.Alive+h.Alive+i.Alive == 3 {
			return 1
		}
	case 0:
		if b.Alive+c.Alive+d.Alive+e.Alive+f.Alive+g.Alive+h.Alive+i.Alive == 3 {
			return 1
		}
	}
	return 0
}

//Seeds the map with new alive lifeforms randomly
func seed(w *[][]lifeform.Lifeform) {
	world := *w
	for i := 0; i < 300; i++ {
		y := randomNumber(5, 25)
		x := randomNumber(5, 25)
		world[y][x].Alive = randomNumber(0, 2)
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
			adjust(&world)
			frame(world)
		}
	}
}

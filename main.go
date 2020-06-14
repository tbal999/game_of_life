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
var cycle = 0

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

//Generates a random number
func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

//Generates a 1 or a 0
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
	fmt.Println("Game of Life - press 'n' for new generated world, 'q' to quit or press enter to run 1 cycle")
	for game == true {
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "n":
			rand.Seed(time.Now().UTC().UnixNano())
			reset(&world)
			seed(&world)
			cycle = 0
		case "q":
			game = false
		default:
			lifeform.Adjust(&world)
			frame(world)
			cycle++
		}
	}
}

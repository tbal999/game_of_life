package lifeform

type Lifeform struct {
	Alive int
	Next  int
}

func Newlife() Lifeform {
	x := Lifeform{0, 0}
	return x
}

func Newline(x int) []Lifeform {
	l := []Lifeform{}
	for i := 0; i <= x; i++ {
		l = append(l, Newlife())
	}
	return l
}

func Newworld(y, x int) [][]Lifeform {
	w := [][]Lifeform{}
	for i := 0; i <= y; i++ {
		w = append(w, Newline(x))
	}
	return w
}

func OnOff(y, x int, w *[][]Lifeform) {
	world := *w
	if world[y][x].Alive == 1 {
		world[y][x].Alive = 0
		world[y][x].Next = 0
	} else if world[y][x].Alive == 0 {
		world[y][x].Alive = 1
		world[y][x].Next = 1
	}
	*w = world
}

//Checks to see if the item in array is one away from the edge to prevent indexing errors
func check(x, y int, z [][]Lifeform) int {
	switcher := ""
	xaxis := x == len(z[0])-1
	yaxis := y == len(z)-1
	isxzero := x == 0
	isyzero := y == 0
	switch xaxis {
	case true:
		switcher += "y"
	case false:
		switcher += "n"
	}
	switch yaxis {
	case true:
		switcher += "y"
	case false:
		switcher += "n"
	}
	switch isxzero {
	case true:
		switcher += "y"
	case false:
		switcher += "n"
	}
	switch isyzero {
	case true:
		switcher += "y"
	case false:
		switcher += "n"
	}
	switch switcher {
	case "nnnn":
		return 0
	case "nnny":
		return 1
	case "nnyn":
		return 2
	case "nnyy":
		return 3
	case "nynn":
		return 4
	case "nyyn":
		return 5
	case "ynnn":
		return 6
	case "ynny":
		return 7
	case "yynn":
		return 8
	}
	return 9 //ignore
}

//Adjusts lifeform state depending on where it is using the 'check' function'
func Adjust(world *[][]Lifeform, l, d []int) {
	i := *world
	for y := range i {
		for x := range i[y] {
			switch check(x, y, i) {
			case 0:
				i[y][x].Next = state(i[y][x],
					i[y-1][x-1],
					i[y-1][x],
					i[y-1][x+1],
					i[y+1][x-1],
					i[y+1][x],
					i[y+1][x+1],
					i[y][x-1],
					i[y][x+1], l, d)
			case 1:
				i[y][x].Next = state(i[y][x],
					i[len(i)-1][x-1],
					i[len(i)-1][x],
					i[len(i)-1][x+1],
					i[y+1][x-1],
					i[y+1][x],
					i[y+1][x+1],
					i[y][x-1],
					i[y][x+1], l, d)
			case 2:
				i[y][x].Next = state(i[y][x],
					i[y-1][len(i[y])-1],
					i[y-1][x],
					i[y-1][x+1],
					i[y+1][len(i[y])-1],
					i[y+1][x],
					i[y+1][x+1],
					i[y][len(i[y])-1],
					i[y][x+1], l, d)
			case 3:
				i[y][x].Next = state(i[y][x],
					i[len(i)-1][len(i[y])-1],
					i[len(i)-1][x],
					i[len(i)-1][x+1],
					i[y+1][len(i[y])-1],
					i[y+1][x],
					i[y+1][x+1],
					i[y][len(i[y])-1],
					i[y][x+1], l, d)
			case 4:
				i[y][x].Next = state(i[y][x],
					i[y-1][x-1],
					i[y-1][x],
					i[y-1][x+1],
					i[0][x-1],
					i[0][x],
					i[0][x+1],
					i[y][x-1],
					i[y][x+1], l, d)
			case 5:
				i[y][x].Next = state(i[y][x],
					i[y-1][len(i[y])-1],
					i[y-1][x],
					i[y-1][x+1],
					i[0][len(i[y])-1],
					i[0][x],
					i[0][x+1],
					i[y][len(i[y])-1],
					i[y][x+1], l, d)
			case 6:
				i[y][x].Next = state(i[y][x],
					i[y-1][x-1],
					i[y-1][x],
					i[y-1][0],
					i[y+1][x-1],
					i[y+1][x],
					i[y+1][0],
					i[y][x-1],
					i[y][0], l, d)
			case 7:
				i[y][x].Next = state(i[y][x],
					i[len(i)-1][x-1],
					i[len(i)-1][x],
					i[len(i)-1][0],
					i[y+1][x-1],
					i[y+1][x],
					i[y+1][0],
					i[y][x-1],
					i[y][0], l, d)
			case 8:
				i[y][x].Next = state(i[y][x],
					i[y-1][x-1],
					i[y-1][x],
					i[y-1][0],
					i[0][x-1],
					i[0][x],
					i[0][0],
					i[y][x-1],
					i[y][0], l, d)
			}
		}
	}
	*world = i
}

//Adjusts the state of the lifeform's next cycle depending on whether they're alive or not.
//You can change this function and get different results.
func state(a, b, c, d, e, f, g, h, i Lifeform, live, dead []int) int {
	total := b.Alive + c.Alive + d.Alive + e.Alive + f.Alive + g.Alive + h.Alive + i.Alive
	switch a.Alive {
	case 1:
		for index := range live {
			if live[index] == total {
				return 1
			}
		}
	case 0:
		for index := range dead {
			if dead[index] == total {
				return 1
			}
		}
	}
	return 0
}

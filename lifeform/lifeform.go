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

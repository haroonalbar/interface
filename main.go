package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	Name   string
	Author string
}

type Count int

func (b Book) String() string {
	return fmt.Sprintf("%s is by %s", b.Name, b.Author)
}

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Log takes in any object that satisfies the interface fmt.Stringer
// here both Book and Count satisfies it cause both have the method 
// String() string which satisfies fmt.Stringer 
func Log(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	b := Book{"Bros before hoes", "Jonny"}
	i := Count(1)

	Log(b)
	Log(i)
}

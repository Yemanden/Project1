package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var s = map[string]Vertex{
	"aaa": {9, 9},
	"hhh": {5, 6},
	"a":   {0, 0},
	"ggg": {4, 4},
}

func main() {
	fmt.Println(s)

}

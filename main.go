// main.go
package main

import (
	"fmt"
)

func double(a float64) float64 {
	return a * 2
}

func main() {
	fmt.Println(double(5))
}

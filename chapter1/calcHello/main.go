package main 

import (
	"fmt"
	"github.com/dshaps10/webdev-with-go/chapter1/calc"
)

func main() {
	var x, y int = 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x,y))
}
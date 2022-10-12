package main

import (
	"fmt"
	"num/numbers"
)

func main() {
	var x int = 10
	fmt.Printf("%d is even a number: %t", x, numbers.Even(x))
}

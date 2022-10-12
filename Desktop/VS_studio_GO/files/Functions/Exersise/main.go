package main

import "fmt"

func f1(a ...int) {
	var s int
	for _, v := range a {
		s += v
	}

	fmt.Println(s)
}

func main() {
	f1()
	f1(1, 2)
}

package main

import "fmt"

func f1() {
	fmt.Println("Calling f1")
}

func f2() {
	fmt.Println("Calling f2")
}

func main() {
	defer f1()
	defer f2()
	f2()
}

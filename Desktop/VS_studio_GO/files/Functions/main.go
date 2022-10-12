package main

import "fmt"

func f1(a ...int) {
	fmt.Printf("%T\n", a)

	fmt.Println("This is a  function")
}
func f2(a ...int) {

	fmt.Println(a)
}
func main() {
	f1()
	f2(69, 78)
}

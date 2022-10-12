package main

import "fmt"

func main() {
	var x int
	var z []int

	fmt.Printf("enter how many numbers you want\n")
	fmt.Scanln(&x)

	sum := 0
	Sum1 := 0

	for i := 0; i < x; i++ {

		var y int
		fmt.Printf("enter the %d number:", i)

		fmt.Scan(&y)
		z = append(z, y)
		sum = sum + y

	}
	for i := 0; i < x; i++ {
		Sum1 = Sum1 + z[i]
	}
	fmt.Printf("the numbers are %v\n", z)
	fmt.Printf("Sum of numbers is %v\n", sum)
	fmt.Printf("Sum of numbers is %v\n", Sum1)
}

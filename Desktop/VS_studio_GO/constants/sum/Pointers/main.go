package main

import "fmt"

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 4
	*price = 500.2
	*name = "Mobile"
	*sold = false
}

func main() {
	quantity, price, name, sold := 4, 243.5, "laptop", true
	fmt.Println(quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changeValuesByPointer():", quantity, price, name, sold)
}

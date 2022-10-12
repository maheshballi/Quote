package main

import (
	"fmt"
	"strings"
)

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

func main() {
	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info)
}

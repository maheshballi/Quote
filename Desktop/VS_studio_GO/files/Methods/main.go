package main

import (
	"fmt"
	"time"
)

type name []string

func (n name) prints() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}
func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := 31 * day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Println(seconds)

	friends := name{"ABD", "VK"}
	friends.prints()
	name.prints(friends)
	friends = append(friends, "mahesh")
	friends.prints()
}

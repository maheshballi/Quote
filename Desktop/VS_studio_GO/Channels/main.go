package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("channel is being sent")

	go send(ch)

	fmt.Println("Channel is being recieved")
	go recieve(ch)

	time.Sleep(time.Second * 1)

}
func send(ch chan int) {
	ch <- 10
}
func recieve(ch chan int) {
	val := <-ch
	fmt.Printf("%d", val)
}

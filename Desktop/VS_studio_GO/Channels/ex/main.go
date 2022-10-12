package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan int) {
	ch <- 1
	fmt.Println("Sending value to channel complete")
}

func receive(ch chan int) {
	//time.Sleep(time.Second * 3)
	fmt.Println("Timeout finished")
	val := <-ch
	fmt.Println(val)

}

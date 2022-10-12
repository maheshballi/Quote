package main

import (
	"fmt"
	"runtime"
	"time"
)

func start() {
	fmt.Println("started")
}
func main() {
	go start()

	fmt.Println("No. of CPUs:", runtime.NumCPU())
	fmt.Println("No. of goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	time.Sleep(time.Second)
	fmt.Println("Arch:", runtime.GOARCH)

}

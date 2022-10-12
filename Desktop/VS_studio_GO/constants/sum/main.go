package main

import (
	"log"
	"os"
)

func main() {
	//var newFile *os.File
	//fmt.Println("%T\n", newFile)

	x, _ := os.Create("demo.txt")
	log.Println(x)

}

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	_, err := os.Create("readFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("b.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

}

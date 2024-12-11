package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the file
	file, err := os.Open("go-101/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err().Error())
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter grade: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	grade, err := strconv.ParseInt(text, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {

		fmt.Println("You are a passing student!")
	} else {
		fmt.Println("You are not a passing student!")
	}
	//fileInfo, _ := os.Stat("go-101/hello_word.go")
	//fmt.Println(fileInfo.Name())

}

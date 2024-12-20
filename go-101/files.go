package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	defer reportPanic()
	scanDirectory("/Users/mac/SWE/go-learn/my_directory")

	//one()
	fmt.Println("Exiting normally")
}
func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

func one() {
	defer fmt.Println("Run one")
	two()
}

func two() {
	defer fmt.Println("Run two")
	three()
}

func three() {
	defer calmDown()
	panic("This call stack's too deep for me")
}

func calmDown() {
	fmt.Println(recover())
}

func scanDirectory(path string) {
	fmt.Println(path)

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fmt.Println("Directory Name ", file.Name())
			if strings.ToLower(file.Name()) == "energy" {
				panic(errors.New("there is an error o")) // cause an error
			}
			scanDirectory(filePath)

		} else {
			fmt.Println("File Name ", file.Name())
		}

	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := scanDirectory("/Users/mac/SWE/go-learn/my_directory")
	if err != nil {
		log.Fatal(err)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fmt.Println("Directory Name ", file.Name())
			err = scanDirectory(filePath)
			if err != nil {
				return err
			}

		} else {
			fmt.Println("File Name ", file.Name())
		}

	}

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Sum is: %0.2f\n", sum)

}

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening file:", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file:", file.Name())
	err := file.Close()
	if err != nil {
		return
	}
}

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	for scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

package main

import "fmt"

func main() {
	accountAge := 2.6
	accountAgeInt := int64(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://www.example.com")
	go responseSize("https://www.flutter.dev")
	go responseSize("https://www.golang.org")
	time.Sleep(2 * time.Second)

	//go a()
	//go b()
	//time.Sleep(time.Second)
	//fmt.Println("end main()")
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	fmt.Print("\n")

}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	fmt.Print("\n")

}

func responseSize(url string) {
	fmt.Println("Getting:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(body))
}

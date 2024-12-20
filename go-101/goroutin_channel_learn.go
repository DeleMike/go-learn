package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	//myChannel := make(chan string)
	//go greeting(myChannel)
	//receivedVal := <-myChannel
	//fmt.Println(receivedVal)

	sizes := make(chan Result)
	urls := []string{
		"https://www.example.com",
		"https://www.flutter.dev",
		"https://www.golang.org",
	}

	for _, url := range urls {
		go responseSize(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)

	}

	//time.Sleep(2 * time.Second)

	//go a()
	//go b()
	//time.Sleep(time.Second)
	//fmt.Println("end main()")

	//channel1 := make(chan string)
	//channel2 := make(chan string)
	//
	//go abc(channel1)
	//go def(channel2)
	//
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//fmt.Println()

	//myChannel := make(chan string)
	//go send(myChannel)
	//reportNap("receiving goroutine", 2)
	//fmt.Println(<-myChannel)
	//fmt.Println(<-myChannel)
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(channel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***Sending value***")
	channel <- "a"
	fmt.Println("***Sending value***")
	channel <- "b"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
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

func responseSize(url string, channel chan Result) {
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

	//fmt.Println(len(body))
	result := Result{
		url, len(body),
	}
	channel <- result

}

type Result struct {
	name string
	size int
}

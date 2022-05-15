package main

import (
	"fmt"
	"sync"
)

func main() {
	/*var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()*/

	//channel
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say2("Bye", c)
	fmt.Println(<-c)
}

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func say2(message string, c chan<- string) {
	c <- message
}

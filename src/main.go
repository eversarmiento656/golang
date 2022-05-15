package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()

}

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

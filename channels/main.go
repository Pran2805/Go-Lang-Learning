package main

import (
	"fmt"
	"sync"
)

// go routines can talk to each other
// we used the channel's
func main() {
	fmt.Println("Channels")

	myChannel := make(chan int, 1) // that one is buffer channel just for error handling
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		if isChannelOpen {

			fmt.Println(<-ch)
			fmt.Println(val)
		}
		wg.Done()
	}(myChannel, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		wg.Done()
	}(myChannel, wg)

	// It will throw an error
	// fmt.Println(<-myChannel)
	// myChannel <- 5

	wg.Wait()
	close(myChannel) // channel should be close below wg
}

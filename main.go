package main

import "fmt"

func main() {
	channel := make(chan int)

	//t2
	go func() {
		channel <- 1 + 1 // <- add new value to channel
	}()

	//t1
	fmt.Println(<-channel) // empty the value channel
	//A thread talk another thread with channels
}

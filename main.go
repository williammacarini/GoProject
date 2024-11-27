package main

import (
	"fmt"
	"time"
)

// main thread
func main() {
	//each sub task here use 2kb in memory
	go counter(5) //two thread
	go counter(5) //three thread
	counter(5)    // run at main thread
}

func counter(count int) {
	for i := range count {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

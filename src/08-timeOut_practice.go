package main

import (
	"fmt"
	"time"
)

func seed(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 10
	// fmt.Println(<-ch)
}
func main() {
	ch := make(chan int)

	go seed(ch)

	select {
	case data := <-ch:
		fmt.Println(data)
	case <-time.After(6 * time.Second):
		fmt.Println("Timed Out...")
	}
}

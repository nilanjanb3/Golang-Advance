package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}

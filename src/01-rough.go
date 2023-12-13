package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	ch <- 10
	close(ch)
	fmt.Println(<-ch)
}

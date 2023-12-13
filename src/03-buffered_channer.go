package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20

	val, _ := <-ch
	fmt.Println(val)
	close(ch)
	val, _ = <-ch
	fmt.Println(val)

	val, _ = <-ch
	fmt.Println(val)
}

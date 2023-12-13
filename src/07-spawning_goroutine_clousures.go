package main

import (
	"fmt"
	"sync"
)

func foo(ch chan int, wg *sync.WaitGroup) {
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	foo(ch, &wg)
	wg.Wait()
	fmt.Println("Done")
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{1, 2, 3, 4, 5}

	sort.Ints(vars)

	fmt.Println(vars)
}

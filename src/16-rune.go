package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := "nilanjan"

	for _, ch := range name {
		fmt.Print(ch, "-->")
		fmt.Print(strconv.QuoteRune(ch), " ")
	}

}

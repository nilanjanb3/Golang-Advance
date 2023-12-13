package main

import (
	"fmt"

	"github.com/nilanjanb3/cryptit/cryptic/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()

	fmt.Println(uuid)

	text := "Hello World"

	fmt.Println(encrypt.Nimbus(text))
}

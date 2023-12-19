package main

import (
	"errors"
	"fmt"
)

func throwError() error {
	return errors.New("This is a custom error")
}

func main() {
	fmt.Println(throwError().Error())
}

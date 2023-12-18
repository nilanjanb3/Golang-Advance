package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	file, err := os.OpenFile("writefile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString("\nHello I'm good")
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

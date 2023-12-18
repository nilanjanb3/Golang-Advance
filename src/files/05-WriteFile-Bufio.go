package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	file, err := os.OpenFile("writefile.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer file.Close()
	writter := bufio.NewWriter(file)

	_, err = writter.WriteString("Hello, I'm good")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	err = writter.Flush()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

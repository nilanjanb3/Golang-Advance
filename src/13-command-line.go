package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	var name string
	var age int

	flag.StringVar(&name, "name", "Default", "Collect Name")
	flag.IntVar(&age, "age", 00, "Collect Age")
	flag.Parse()
	/*
		go run 13-command-line.go -name Nilanjan -age 24
		If we pass any other argument before the mentioned one, it'll not work
	*/
	fmt.Printf("Hi %v, you are %v", name, age)
}

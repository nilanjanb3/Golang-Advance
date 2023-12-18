package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Split(scanner.Text(), " ")
	age, _ := strconv.Atoi(str[2])
	fmt.Printf("Hi %v %v, Your age is %v", str[0], str[1], age)

}

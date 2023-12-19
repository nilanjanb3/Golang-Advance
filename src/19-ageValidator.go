package main

import (
	"fmt"
)

// CustomError is a custom error type
type CustomAgeValidationError struct {
	Message string
}

// Error returns the error message
func (e *CustomAgeValidationError) Error() string {
	return e.Message
}

// Function that may return a custom error
func isEligable(age int) (bool, error) {
	// Some condition that triggers the custom error
	if age < 18 {
		return false, &CustomAgeValidationError{Message: "Age is not valid"}
	} else {
		return true, nil
	}
}

func main() {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanf("%d", &age)
	flag, err := isEligable(age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(flag)
	}
}

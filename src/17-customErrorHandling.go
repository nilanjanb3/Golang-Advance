package main

import (
	"fmt"
)

// CustomError is a custom error type
type CustomError struct {
	Message string
}

// Error returns the error message
func (e *CustomError) Error() string {
	return e.Message
}

// Function that may return a custom error
func myFunction() error {
	// Some condition that triggers the custom error
	if true {
		return &CustomError{Message: "This is a custom error."}
	}

	// No error occurred
	return nil
}

func main() {
	err := myFunction()
	if err != nil {
		// Check if the error is of the custom type
		if customErr, ok := err.(*CustomError); ok {
			fmt.Printf("Custom error: %s\n", customErr.Message)
		} else {
			// Handle other types of errors
			fmt.Printf("Error: %s\n", err.Error())
		}
	} else {
		fmt.Println("No error occurred.")
	}
}

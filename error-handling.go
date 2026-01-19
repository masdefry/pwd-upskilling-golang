package main

import (
	"errors"
	"fmt"
)

func checkNumber(num int) (string, error) {
	if num < 0 {
		// Return an error when the number is negative
		return "", errors.New("number is negative")
	}
	// Return a success message and nil for no error
	return "number is positive", nil
}

func main() {
	// Call the function and check the error
	result, err := checkNumber(-5)
	if err != nil {
		// If an error is returned, handle it
		fmt.Println("Error:", err)
	} else {
		// If no error, print the result
		fmt.Println(result)
	}
}

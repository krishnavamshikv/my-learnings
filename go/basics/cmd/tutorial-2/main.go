package main

import (
	"errors"
	"fmt"
)

// Functions in go.

func main() {
	// printMe("hello krishna")
	var numerator int = 10
	var denominator int = 0
	var result, remainder, err = divide(numerator, denominator)

	// some modules might come with err built in , this is the standard way of handling errors in go.
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The integer division is %v and modulo is %v \n", result, remainder)
	}
}

// func printMe(str string) {
// 	fmt.Println(str)
// }

func divide(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}

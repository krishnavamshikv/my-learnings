package main

import "fmt"

// strings , runes and bytes

// strings are stored in utf-8 format

func main() {
	var myString = []rune("résumé") // need to learn more on runes
	var indexed = myString[0]
	fmt.Println(myString, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

}

package main

import "fmt"

// data types in go

func main() {
	var myInt int = 24
	fmt.Println(myInt)

	inferedInt := 25
	fmt.Println(inferedInt)

	myInt = 28

	var myFun int = 24
	fmt.Println(myFun)

	var eightBit uint8 = 255
	fmt.Println(eightBit)

	var myString string = "hello world"
	fmt.Println(myString)

	var mySeconSting string
	fmt.Println(mySeconSting)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	v1, v2 := 1, 3
	fmt.Println(v1 + v2)

	const myConst string = "this is a const"

}

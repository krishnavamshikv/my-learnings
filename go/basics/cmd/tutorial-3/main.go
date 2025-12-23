package main

import "fmt"

// Arrays , loops and slices

func main() {
	// Arrays are fixed length
	var intArr [3]int32 = [...]int32{2, 3, 4}
	fmt.Println(intArr)

	intArr2 := [3]int32{5, 6, 7}
	fmt.Println(intArr2)

	// Slices are like vectors
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length of slice is %v and capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length of slice is %v and capacity is %v\n", len(intSlice), cap(intSlice))

}

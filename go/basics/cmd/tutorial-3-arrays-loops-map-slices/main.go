package main

import (
	"fmt"
	"time"
)

// Arrays , loops , map and slices

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
	var intSlice2 []int32 = []int32{5, 6, 7}
	intSlice2 = append(intSlice, intSlice2...) // appending an slice using spread method
	fmt.Println(intSlice2)

	var intSlice3 []int32 = make([]int32, 3, 10) // using make to create slice make(type , length , capacity)
	fmt.Println(intSlice3)

	// map

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 24, "krishna": 23}
	fmt.Println(myMap2["krishna"])

	// if the value exists in map or not , if we enter value that is not map it will return default value

	var age, ok = myMap2["adam"] // map returns the value true if element is present else false.
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Invalid name.")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v  age: %v \n", name, age)
	}

	for i, v := range intSlice2 {
		fmt.Printf("Index: %v value: %v \n", i, v)
	}

	// speed test of setting capacity ahead of time and not setting
	// try to declare slices with capacity
	var n int = 1000000
	var testSlice = []int{} // setting without declaring capacity
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n)) // this took 48.12 ms
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))   // this took 10.80 ms 🤯

}

// function to calculate time duration to count n times
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

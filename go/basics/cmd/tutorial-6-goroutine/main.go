package main

import (
	"fmt"
	"sync"
	"time"
)

// without go routines below code will take 5 seconds to complete , but with routines it will take only one second 🤯
// because of parallel execution , tasks are allocated to multiple cores (if there is only single core it will take 5 seconds)
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := range len(dbData) {
		wg.Add(1) // this will increase the counter everytime before calling go routine
		go dbCall(i)
	}
	wg.Wait() // this will wait till the counter becomes zero
	fmt.Printf("\nTime of execution: %v\n", time.Since(t0))
}

func dbCall(i int) {
	// simulate db delay
	var delay float32 = 1000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	wg.Done()
}

// we shouldn't read and write using goroutines directly that will corrupt the data instead we have to use sync.Mutex.

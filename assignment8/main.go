package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var m sync.Mutex

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	// m.Lock()
	n := 0
	// m.Unlock()
	wg.Add(1)
	go func() {
		nIsEven := isEven(n)

		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		wg.Done()
	}()

	go func() {
		wg.Wait()
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}

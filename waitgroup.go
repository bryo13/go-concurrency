package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, n int) {
		defer wg.Done()
		fmt.Printf("hello %v times!\n", n)
	}

	greetingCount := 5
	var wg sync.WaitGroup

	wg.Add(greetingCount) // notice where we add the Add() function

	for i := 0; i < greetingCount; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}

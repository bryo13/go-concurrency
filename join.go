package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salut := range []string{"hola", "jambo", "hi"} {
		wg.Add(1)
		go func(salut string) {
			defer wg.Done()
			fmt.Println(salut)
		}(salut)
	}

	wg.Wait() // join point
}

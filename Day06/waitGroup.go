package main

import (
	"fmt"
	"sync"
)

func main() {
	panicAndRecover()

	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for _, v := range arr {
				fmt.Println(i, v)
			}
		}(i)
	}

	fmt.Println("Printing before waiting")
	wg.Wait()
	fmt.Println("Printing after waiting")
}

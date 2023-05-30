package main

import (
	"fmt"
	"sync"
	"time"
)

var product int = 1

func prod(i, e int) {
	if i%2 == 0 {
		time.Sleep(time.Millisecond * 2000)
	}
	product *= e
}

type mutexes struct {
	ch chan int
	newCh chan int
	mux *sync.Mutex
}

func practice() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// var ch = make(chan int)
	var sum int

	
	var obj = new(mutexes)
	obj.ch = make(chan int)
	obj.mux = &sync.Mutex{}
	
	for i, element := range arr {
		go prod(i, element)
		// fmt.Println(product)
	}

	go func(){
		obj.mux.Lock()
		defer obj.mux.Unlock()
		for _, element := range arr {
			obj.ch <- element
		}

		close(obj.ch);
	}()

	go func() {
		for true {
			element, ok := <- obj.ch
			if !ok {
				break
			}

			fmt.Println(element, sum)
			sum += element
		}
	}()
	
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("Final value of product: ", product)
	fmt.Println("Final value of sum: ", sum)


	obj.newCh = make(chan int)
	// var newCh = make(chan int)
	go obj.fibonacci(90);
	for e := range obj.newCh {
		fmt.Println("fibonacci: ", e);
	}

}

func (s *mutexes)fibonacci(n int) {
	for i, x, y := 0, 0, 1; i <= n; i++ {
		s.newCh <- x
		x, y = y, x + y
	} 

	close(s.newCh);
}
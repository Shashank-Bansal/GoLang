package main

import (
	"fmt"
)

func panicAndRecover() {
	defer func() {
		if recover() != nil {
			fmt.Println("Trying to recover")
		}
	}()

	panic("panicking")
}

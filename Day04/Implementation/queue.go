package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr    []int
	length int
}

func (q *Queue) add(element int) {
	if q.isEmpty() {
		q.arr = make([]int, 0)
	}

	q.arr = append(q.arr, element)
	q.length++
}

func (q *Queue) remove() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("Queue is empty")
	}

	pop := q.arr[0]
	q.arr = q.arr[1:]
	q.length--
	return pop, nil
}

func (q *Queue) peek() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("Queue is empty")
	}

	peek := q.arr[0]
	return peek, nil
}

func (q *Queue) size() int {
	return q.length
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

func (q *Queue) print() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	for _, element := range q.arr {
		fmt.Printf("%d, ", element)
	}

	fmt.Println()
}

// func main() {
// 	var queue = new(Queue)
// 	fmt.Println("\nCreating a new Queue...")

// 	for true {
// 		var choice int
// 		fmt.Println("\nEnter 1 to add a new element in the Queue")
// 		fmt.Println("Enter 2 to remove element from the Queue")
// 		fmt.Println("Enter 3 to get peek element from the Queue")
// 		fmt.Println("Enter 4 to print Queue")
// 		fmt.Println("Enter 5 to print length of Queue")
// 		fmt.Println("Enter 6 to check Queue is empty or not")
// 		fmt.Println("Enter 0 to exit")
// 		fmt.Print("Enter your choice: ")

// 		_, err := fmt.Scan(&choice)
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			continue
// 		}

// 		fmt.Println()

// 		if choice == 0 {
// 			break
// 		} else if choice == 1 {
// 			var element int
// 			fmt.Print("Enter the element you want to add: ")
// 			_, err := fmt.Scan(&element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			queue.add(element)
// 			fmt.Printf("Element %d added successfully", element)
// 		} else if choice == 2 {
// 			element, err := queue.remove()

// 			if err != nil {
// 				fmt.Println(err.Error())
// 			} else {
// 				fmt.Printf("The removed element is %d\n", element)
// 			}
// 		} else if choice == 3 {
// 			element, err := queue.peek()

// 			if err != nil {
// 				fmt.Println(err.Error())
// 			} else {
// 				fmt.Printf("The peek element is %d\n", element)
// 			}
// 		} else if choice == 4 {
// 			queue.print()
// 		} else if choice == 5 {
// 			fmt.Println("The length of the queue is", queue.length)
// 		} else if choice == 6 {
// 			if queue.isEmpty() {
// 				fmt.Println("Queue is empty")
// 			} else {
// 				fmt.Println("Queue is not empty")
// 			}
// 		}
// 	}
// }

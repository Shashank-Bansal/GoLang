package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	arr    []int
	length int
}

func (s *Stack) push(element int) {
	if s.isEmpty() {
		s.arr = make([]int, 0)
	}

	s.arr = append(s.arr, element)
	s.length++
}

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return -1, errors.New("Stack is empty")
	}

	pop := s.arr[s.length-1]
	s.arr = s.arr[:s.length-1]
	s.length--
	return pop, nil
}

func (s *Stack) peek() (int, error) {
	if s.isEmpty() {
		return -1, errors.New("Stack is empty")
	}

	peek := s.arr[s.length-1]
	return peek, nil
}

func (s *Stack) size() int {
	return s.length
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

func (s *Stack) print() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	for _, element := range s.arr {
		fmt.Printf("%d, ", element)
	}

	fmt.Println()
}

func main() {
	var stack = new(Stack)
	fmt.Println("\nCreating a new Stack...")

	for true {
		var choice int
		fmt.Println("\nEnter 1 to push a new element in the stack")
		fmt.Println("Enter 2 to pop element from the stack")
		fmt.Println("Enter 3 to get peek element from the stack")
		fmt.Println("Enter 4 to print stack")
		fmt.Println("Enter 5 to print length of stack")
		fmt.Println("Enter 6 to check stack is empty or not")
		fmt.Println("Enter 0 to exit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		fmt.Println()

		if choice == 0 {
			break
		} else if choice == 1 {
			var element int
			fmt.Print("Enter the element you want to add: ")
			_, err := fmt.Scan(&element)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			stack.push(element)
			fmt.Printf("Element %d added successfully", element)
		} else if choice == 2 {
			element, err := stack.pop()

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The popped element is %d\n", element)
			}
		} else if choice == 3 {
			element, err := stack.peek()

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The peek element is %d\n", element)
			}
		} else if choice == 4 {
			stack.print()
		} else if choice == 5 {
			fmt.Println("The length of the stack is", stack.length)
		} else if choice == 6 {
			if stack.isEmpty() {
				fmt.Println("Stack is empty")
			} else {
				fmt.Println("Stack is not empty")
			}
		}
	}
}

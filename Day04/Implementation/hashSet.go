package main

import "fmt"

type HashSet struct {
	arr    map[int]bool
	length int
}

func (h *HashSet) add(element int) bool {
	if h.length == 0 {
		h.arr = make(map[int]bool)
	}

	if h.contains(element) {
		return false
	}

	h.arr[element] = true
	h.length++
	return true
}

func (h *HashSet) contains(element int) bool {
	for key := range h.arr {
		if element == key {
			return true
		}
	}

	return false
}

func (h *HashSet) remove(element int) bool {
	if !h.contains(element) {
		return false
	}

	delete(h.arr, element)
	h.length--
	return true
}

func (h *HashSet) print() {
	if h.length == 0 {
		fmt.Println("The HashSet is empty")
		return
	}

	for key := range h.arr {
		fmt.Printf("%d, ", key)
	}
}

// func main() {
// 	var hashSet = new(HashSet)
// 	fmt.Println("\nCreating a new HashSet...")

// 	for true {
// 		var choice int
// 		fmt.Println("\nEnter 1 to add a new element")
// 		fmt.Println("Enter 2 to remove a element")
// 		fmt.Println("Enter 3 to check if element exists")
// 		fmt.Println("Enter 4 to print hashSet")
// 		fmt.Println("Enter 5 to print length of hashSet")
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

// 			if hashSet.add(element) {
// 				fmt.Printf("Element %d added successfully", element)
// 			} else {
// 				fmt.Printf("Element %d already exists in hashSet", element)
// 			}
// 		} else if choice == 2 {
// 			var element int
// 			fmt.Print("Enter the element you want to remove: ")
// 			_, err := fmt.Scan(&element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			if hashSet.remove(element) {
// 				fmt.Printf("Element %d removed successfully", element)
// 			} else {
// 				fmt.Printf("Element %d doesn't exists in hashSet", element)
// 			}
// 		} else if choice == 3 {
// 			var element int
// 			fmt.Print("Enter the element you want to check: ")
// 			_, err := fmt.Scan(&element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			if hashSet.contains(element) {
// 				fmt.Printf("Element %d exists in hashSet\n", element)
// 			} else {
// 				fmt.Printf("Element %d doesn't exists in hashSet\n", element)
// 			}
// 		} else if choice == 4 {
// 			hashSet.print()
// 		} else if choice == 5 {
// 			fmt.Println("The length of the hashSet is", hashSet.length)
// 		}
// 	}
// }

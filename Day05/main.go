package main

import "fmt"
// import LL "./Implementation/linkedList"

func main() {
	fmt.Printf("\nCreating a new LinkedList...\n")
	// var linkedList LinkedList[float64]
	var linkedList LL.LinkedList[float64]

	for true {
		var choice int
		fmt.Println("\nEnter 1 to add a new node at last")
		fmt.Println("Enter 2 to delete last node")
		fmt.Println("Enter 3 to add a new node at start")
		fmt.Println("Enter 4 to delete start node")
		fmt.Println("Enter 5 to add a new node at particular index")
		fmt.Println("Enter 6 to delete a node at particular index")
		fmt.Println("Enter 7 to update value of a node at particular index")
		fmt.Println("Enter 8 to print linked list")
		fmt.Println("Enter 9 to print length of linked list")
		fmt.Println("Enter 0 to exit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if choice == 0 {
			break
		} else if choice == 1 {
			var element float64
			fmt.Print("\nEnter the element you want to add: ")
			_, err := fmt.Scan(&element)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			linkedList.add(element)
		} else if choice == 2 {
			if linkedList.delete() {
				fmt.Println("Deleted successfully")
			} else {
				fmt.Println("unable to delete")
			}
		} else if choice == 3 {
			var element float64
			fmt.Print("\nEnter the element you want to add: ")
			_, err := fmt.Scan(&element)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			linkedList.addAtHead(element)
		} else if choice == 4 {
			if linkedList.deleteAtHead() {
				fmt.Println("Deleted successfully")
			} else {
				fmt.Println("unable to delete")
			}
		} else if choice == 5 {
			var element float64
			var index int
			fmt.Print("\nEnter the index and element you want to add: ")
			_, err := fmt.Scan(&index, &element)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if linkedList.insert(index, element) {
				fmt.Printf("Element %v added successfully at index %v", element, index)
			} else {
				fmt.Printf("unable to add element %v at index %v because index is either negative number or greater than the length of linkedList", element, index)
			}
		} else if choice == 6 {
			var index int
			fmt.Print("\nEnter the index of node you want to delete: ")
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if linkedList.deleteAtIndex(index) {
				fmt.Printf("Node of index %v deleted successfully", index)
			} else {
				fmt.Printf("unable to delete node of index %v because index is either negative number or greater than the length of linkedList", index)
			}
		} else if choice == 7 {
			var value float64
			var index int
			fmt.Print("\nEnter the index and new value of node you want to update: ")
			_, err := fmt.Scan(&index, &value)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if linkedList.update(index, value) {
				fmt.Printf("Node updated successfully at index %v", index)
			} else {
				fmt.Printf("unable to update at index %v because index is either negative number or greater than the length of linkedList", index)
			}
		} else if choice == 8 {
			linkedList.print()
		} else if choice == 9 {
			fmt.Println("The length of linkedList is", linkedList.length)
		}
	}
}

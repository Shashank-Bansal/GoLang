package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) add(element int) {
	if l.length == 0 {
		l.head = &Node{val: element}
	} else {
		node := l.lastNode()
		node.next = &Node{val: element}
	}

	l.length++
}

func (l *LinkedList) delete() bool {
	if l.length == 0 {
		fmt.Println("LinkedList is empty")
		return false
	} else if l.length == 1 {
		l.head = nil
	} else {
		node := l.head
		for node.next != nil && node.next.next != nil {
			node = node.next
		}
		node.next = nil
	}

	l.length--
	return true
}

func (l *LinkedList) lastNode() *Node {
	node := l.head
	for node.next != nil {
		node = node.next
	}

	return node
}

func (l *LinkedList) print() {
	if l.length == 0 {
		fmt.Println("The linkedList is empty")
		return
	}

	node := l.head
	for node != nil {
		fmt.Print(node.val, " -> ")
		node = node.next
	}

	fmt.Println()
}

func (l *LinkedList) addAtHead(element int) {
	newHead := &Node{val: element}
	newHead.next = l.head
	l.head = newHead
	l.length++
}

func (l *LinkedList) deleteAtHead() bool {
	if l.length == 0 {
		fmt.Println("LinkedList is empty")
		return false
	}

	l.head = l.head.next
	l.length--
	return true
}

func (l *LinkedList) insert(index, element int) bool {
	if index < 0 || index > l.length {
		return false
	}

	if index == 0 {
		l.addAtHead(element)
	} else if l.length == index {
		l.add(element)
	} else {
		node := l.findNodeAtIndex(index - 1)
		newNode := &Node{val: element}
		newNode.next = node.next
		node.next = newNode
		l.length++
	}

	return true
}

func (l *LinkedList) deleteAtIndex(index int) bool {
	if index < 0 || index > l.length {
		return false
	}

	if index == 0 {
		return l.deleteAtHead()
	} else if l.length == index {
		return l.delete()
	} else {
		node := l.findNodeAtIndex(index - 1)
		node.next = node.next.next
		l.length--
	}

	return true
}

func (l *LinkedList) findNodeAtIndex(index int) *Node {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *LinkedList) update(index, value int) bool {
	if index < 0 || index > l.length {
		return false
	}

	node := l.findNodeAtIndex(index)
	node.val = value
	return true
}

func main() {
	var linkedList = new(LinkedList)
	fmt.Println("\nCreating a new LinkedList...")

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
			var element int
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
			var element int
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
			var element, index int
			fmt.Print("\nEnter the index and element you want to add: ")
			_, err := fmt.Scan(&index, &element)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if linkedList.insert(index, element) {
				fmt.Printf("Element %d added successfully at index %d", element, index)
			} else {
				fmt.Printf("unable to add element %d at index %d because index is either negative number or greater than the length of linkedList", element, index)
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
				fmt.Printf("Node of index %d deleted successfully", index)
			} else {
				fmt.Printf("unable to delete node of index %d because index is either negative number or greater than the length of linkedList", index)
			}
		} else if choice == 7 {
			var value, index int
			fmt.Print("\nEnter the index and new value of node you want to update: ")
			_, err := fmt.Scan(&index, &value)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if linkedList.update(index, value) {
				fmt.Printf("Node updated successfully at index %d", index)
			} else {
				fmt.Printf("unable to update at index %d because index is either negative number or greater than the length of linkedList", index)
			}
		} else if choice == 8 {
			linkedList.print()
		} else if choice == 9 {
			fmt.Println("The length of linkedList is", linkedList.length)
		}
	}
}

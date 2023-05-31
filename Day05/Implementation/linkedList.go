package implementation

import "fmt"
// import 	"reflect"

type Node[T any] struct {
	val  T
	next *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func (l *LinkedList[T]) Add(element T) {
	if l.length == 0 {
		l.head = &Node[T]{val: element}
	} else {
		node := l.LastNode()
		node.next = &Node[T]{val: element}
	}

	l.length++
}

func (l *LinkedList[T]) Delete() bool {
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

func (l *LinkedList[T]) LastNode() *Node[T] {
	node := l.head
	for node.next != nil {
		node = node.next
	}

	return node
}

func (l *LinkedList[T]) Print() {
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

func (l *LinkedList[T]) AddAtHead(element T) {
	newHead := &Node[T]{val: element}
	newHead.next = l.head
	l.head = newHead
	l.length++
}

func (l *LinkedList[T]) DeleteAtHead() bool {
	if l.length == 0 {
		fmt.Println("LinkedList is empty")
		return false
	}

	l.head = l.head.next
	l.length--
	return true
}

func (l *LinkedList[T]) Insert(index int, element T) bool {
	if index < 0 || index > l.length {
		return false
	}

	if index == 0 {
		l.AddAtHead(element)
	} else if l.length == index {
		l.Add(element)
	} else {
		node := l.FindNodeAtIndex(index - 1)
		newNode := &Node[T]{val: element}
		newNode.next = node.next
		node.next = newNode
		l.length++
	}

	return true
}

func (l *LinkedList[T]) DeleteAtIndex(index int) bool {
	if index < 0 || index > l.length {
		return false
	}

	if index == 0 {
		return l.DeleteAtHead()
	} else if l.length == index {
		return l.Delete()
	} else {
		node := l.FindNodeAtIndex(index - 1)
		node.next = node.next.next
		l.length--
	}

	return true
}

func (l *LinkedList[T]) FindNodeAtIndex(index int) *Node[T] {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *LinkedList[T]) Update(index int, value T) bool {
	if index < 0 || index > l.length {
		return false
	}

	node := l.FindNodeAtIndex(index)
	node.val = value
	return true
}

func (l LinkedList[T]) Length() int {
	return l.length;
}


// func getTypeByName(typeName string) (reflect.Type, error) {
// 	switch typeName {
// 	case "int":
// 		return reflect.TypeOf(0), nil
// 	case "string":
// 		return reflect.TypeOf(""), nil
// 	case "bool":
// 		return reflect.TypeOf(false), nil
// 	case "float64":
// 		return reflect.TypeOf(float64(0)), nil
// 	default:
// 		return nil, fmt.Errorf("unknown type: %s", typeName)
// 	}
// }




// func main() {
// // start:
// 	// var choice int
// 	// fmt.Println("\nEnter 1 to create linkedList of int")
// 	// fmt.Println("Enter 2 to create linkedList of string")
// 	// fmt.Println("Enter 3 to create linkedList of bool")
// 	// fmt.Println("Enter 4 to create linkedList of float")
// 	// _, err := fmt.Scan(&choice)
// 	// if err != nil {
// 	// 	fmt.Println("Error reading input:", err)
// 	// 	goto start
// 	// }

// 	// var dataType string
// 	// switch choice {
// 	// 	case 1:
// 	// 		dataType = "int"
// 	// 	case 2:
// 	// 		dataType = "string"
// 	// 	case 3:
// 	// 		dataType = "bool"
// 	// 	case 4:
// 	// 		dataType = "float64"
// 	// 	default:
// 	// 		goto start
// 	// }
// 	// t, err := getTypeByName(dataType)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	fmt.Printf("\nCreating a new LinkedList...\n")
// 	// var linkedList LinkedList[reflect.New(t).Elem()]
// 	var linkedList LinkedList[float64]
// 	// var linkedList LinkedList[float64]
// 	// var linkedList LinkedList[string]
// 	// var linkedList LinkedList[bool]
	
// 	// // linkedList := reflect.New(reflect.SliceOf(t)).Elem()
// 	// fmt.Println(reflect.New(reflect.SliceOf(t)).Elem())
// 	// fmt.Println(t)
// 	// var s t
// 	// s := 
// 	// fmt.Println(reflect.TypeOf(t))


// 	for true {
// 		var choice int
// 		fmt.Println("\nEnter 1 to add a new node at last")
// 		fmt.Println("Enter 2 to delete last node")
// 		fmt.Println("Enter 3 to add a new node at start")
// 		fmt.Println("Enter 4 to delete start node")
// 		fmt.Println("Enter 5 to add a new node at particular index")
// 		fmt.Println("Enter 6 to delete a node at particular index")
// 		fmt.Println("Enter 7 to update value of a node at particular index")
// 		fmt.Println("Enter 8 to print linked list")
// 		fmt.Println("Enter 9 to print length of linked list")
// 		fmt.Println("Enter 0 to exit")
// 		fmt.Print("Enter your choice: ")

// 		_, err := fmt.Scan(&choice)
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			continue
// 		}
// 		if choice == 0 {
// 			break
// 		} else if choice == 1 {
// 			var element float64
// 			fmt.Print("\nEnter the element you want to add: ")
// 			_, err := fmt.Scan(&element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}
// 			linkedList.Add(element)
// 		} else if choice == 2 {
// 			if linkedList.Delete() {
// 				fmt.Println("Deleted successfully")
// 			} else {
// 				fmt.Println("unable to delete")
// 			}
// 		} else if choice == 3 {
// 			var element float64
// 			fmt.Print("\nEnter the element you want to add: ")
// 			_, err := fmt.Scan(&element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}
// 			linkedList.AddAtHead(element)
// 		} else if choice == 4 {
// 			if linkedList.DeleteAtHead() {
// 				fmt.Println("Deleted successfully")
// 			} else {
// 				fmt.Println("unable to delete")
// 			}
// 		} else if choice == 5 {
// 			var element float64
// 			var index int
// 			fmt.Print("\nEnter the index and element you want to add: ")
// 			_, err := fmt.Scan(&index, &element)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			if linkedList.Insert(index, element) {
// 				fmt.Printf("Element %v added successfully at index %v", element, index)
// 			} else {
// 				fmt.Printf("unable to add element %v at index %v because index is either negative number or greater than the length of linkedList", element, index)
// 			}
// 		} else if choice == 6 {
// 			var index int
// 			fmt.Print("\nEnter the index of node you want to delete: ")
// 			_, err := fmt.Scan(&index)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			if linkedList.DeleteAtIndex(index) {
// 				fmt.Printf("Node of index %v deleted successfully", index)
// 			} else {
// 				fmt.Printf("unable to delete node of index %v because index is either negative number or greater than the length of linkedList", index)
// 			}
// 		} else if choice == 7 {
// 			var value float64
// 			var index int
// 			fmt.Print("\nEnter the index and new value of node you want to update: ")
// 			_, err := fmt.Scan(&index, &value)
// 			if err != nil {
// 				fmt.Println("Error reading input:", err)
// 				continue
// 			}

// 			if linkedList.Update(index, value) {
// 				fmt.Printf("Node updated successfully at index %v", index)
// 			} else {
// 				fmt.Printf("unable to update at index %v because index is either negative number or greater than the length of linkedList", index)
// 			}
// 		} else if choice == 8 {
// 			linkedList.Print()
// 		} else if choice == 9 {
// 			fmt.Println("The length of linkedList is", linkedList.length)
// 		}
// 	}
// }

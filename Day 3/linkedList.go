package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func (n *Node) add (val int) {
	n.next = &Node{val: val}
}

type List struct {
	head  *Node
	// next *Node
}

func main() {
	// var head Node

	// node := &head
	// Head := &head

	// for true {
	// 	var v int
	// 	fmt.Print("Enter a number or 0 to exit: ")
	// 	_, err := fmt.Scan(&v)
	// 	if err != nil {
	// 		fmt.Println("Error reading input:", err)
	// 		return
	// 	}
	// 	if (v == 0) {
	// 		break;
	// 	}

	// 	node.val = v
	// 	var nextNode Node
	// 	node.next = &nextNode
	// 	node = node.next
	// }

	// node = Head
	// for node != nil {
	// 	fmt.Println(node.val) 
	// 	node = node.next
	// }


	var list List
	list.head = new (Node)

	node := list.head

	for true {
		var v int
		fmt.Print("Enter a number or 0 to exit: ")
		_, err := fmt.Scan(&v)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if (v == 0) {
			break;
		}

		node.add(v);
		node = node.next
	}

	node = list.head
	for node != nil {
		fmt.Println(node.val) 
		node = node.next
	}
}

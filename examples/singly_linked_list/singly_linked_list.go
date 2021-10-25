package main

import (
	"fmt"

	singlylinkedlist "github.com/gamberooni/go-dsa/singly_linked_list"
)

// stack can be used to reverse numbers
func main() {
	l := singlylinkedlist.SinglyLinkedList{}
	l.Insert(6)
	l.Insert(5)
	l.Insert(4)

	l.Display()              // shows how the linked list looks like
	fmt.Println(l.Search(5)) // returns 1 (index of the node)

	l.InsertAt(1, 7) // linked list becomes 6, 7, 5, 4
	l.Display()

	l.DeleteAt(0) // linked list becomes 7, 5, 4
	l.Display()

	l.Delete() // linked list becomes 7, 5
	l.Display()

	l.Insert(8) // linked list becomes 7, 5, 8
	l.Display()

	l.Sort() // sort the linked list in ascending order
	l.Display()
}

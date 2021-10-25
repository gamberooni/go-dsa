package main

import (
	doublylinkedlist "github.com/gamberooni/go-dsa/doubly_linked_list"
)

// stack can be used to reverse numbers
func main() {
	l := doublylinkedlist.DoublyLinkedList{}
	l.Insert(6)
	l.Insert(5)
	l.Insert(4)

	l.Display() // shows how the linked list looks like

	l.InsertAt(1, 7) // linked list becomes 6, 7, 5, 4
	l.Display()

	l.DeleteAt(0) // linked list becomes 7, 5, 4
	l.Display()

	l.Delete() // linked list becomes 7, 5
	l.Display()

	l.Insert(8) // linked list becomes 7, 5, 8
	l.Display()
}

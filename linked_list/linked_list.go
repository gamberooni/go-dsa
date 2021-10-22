package linkedlist

import "fmt"

const (
	ErrOutOfBounds = LinkedListError("specified index out of bounds")
)

type LinkedListError string

func (l LinkedListError) Error() string {
	return string(l)
}

type Node struct {
	value int
	next  *Node // next pointer of the last node is nil
}

type LinkedList struct {
	head   *Node // points to the first node of the linked list
	length int
}

// GetNode traverses to the node of the specified index and returns that node
func (l *LinkedList) GetNode(index int) (*Node, error) {
	if (index > l.length-1) || index < 0 {
		return nil, ErrOutOfBounds
	}

	ptr := l.head // initialize a pointer that points to the first node in the linked list
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	return ptr, nil
}

// Insert inserts new node at the end of the linked list
func (l *LinkedList) Insert(value int) {
	node := Node{}
	node.value = value
	// if linked list is empty
	if l.length == 0 {
		l.head = &node
		l.length++
		return
	}
	// if not empty, traverse to the end and add the node
	ptr, _ := l.GetNode(l.length - 1)
	ptr.next = &node
	l.length++
}

// InsertAt inserts a node at the specified index
func (l *LinkedList) InsertAt(index, value int) error {
	// traverse to the specified index
	// point the previous node of the specified index to the new node
	// point the new node to the node that was originally at the specified index
	newNode := Node{}
	newNode.value = value
	originalNode, err := l.GetNode(index)
	if err != nil {
		return err
	}
	previousNode, err := l.GetNode(index - 1)
	if err != nil {
		return err
	}
	previousNode.next = &newNode
	newNode.next = originalNode
	l.length++
	return nil
}

// Delete deletes the last node of the linked list
func (l *LinkedList) Delete() error {
	ptr, err := l.GetNode(l.length - 2)
	if err != nil {
		return err
	}
	ptr.next = nil
	l.length--
	return nil
}

// DeleteAt deletes the node at the specified index
func (l *LinkedList) DeleteAt(index int) error {
	if index == 0 {
		l.head = l.head.next
		l.length--
		return nil
	}

	ptr, err := l.GetNode(index - 1)
	if err != nil {
		return err
	}
	ptr.next, err = l.GetNode(index + 1)
	if err != nil {
		return err
	}
	l.length--
	return nil
}

// Search searches for the specified value from the linked list and returns the index of the node
// returns -1 if not found or if linked list is empty
func (l *LinkedList) Search(value int) int {
	if l.length == 0 {
		return -1
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if value == ptr.value {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

func (l *LinkedList) Sort() {
	if l.length == 0 {
		return
	}

	// use bubble sort
	currentNode := l.head
	var nextNode *Node
	for currentNode.next != nil { // while not the end of the linked list
		nextNode = currentNode.next

		for nextNode != nil { // while next node is not the last node
			if currentNode.value > nextNode.value {
				// swap values
				currentNode.value, nextNode.value = nextNode.value, currentNode.value
			}
			// finished swapping
			// assign the next node variable to the next node
			nextNode = nextNode.next
		}
		// assign the current node variable to the next node
		currentNode = currentNode.next
	}
}

// Display shows the nodes in the linked list
func (l *LinkedList) Display() {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d--->", ptr.value)
		ptr = ptr.next
	}
	fmt.Print("\n")
}

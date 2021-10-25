package doublylinkedlist

import "fmt"

const (
	ErrOutOfBounds     = DoublyLinkedListError("index out of bounds")
	ErrEmptyLinkedList = DoublyLinkedListError("cannot delete from empyty linked list")
)

type DoublyLinkedListError string

func (e DoublyLinkedListError) Error() string {
	return string(e)
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Insert inserts a node at the end of the linked list
func (l *DoublyLinkedList) Insert(value int) {
	node := Node{}
	node.value = value

	// if linked list is empty
	if l.length == 0 {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	lastNode, _ := l.GetNode(l.length - 1) // get the last node
	lastNode.next = &node
	node.prev = lastNode
	l.tail = &node

	l.length++
}

func (l *DoublyLinkedList) InsertAt(index, value int) error {
	node := Node{}
	node.value = value

	if index == l.length-1 { // if the specified index is the last node
		l.Insert(value)
	} else {
		previousNode, err := l.GetNode(index - 1)
		if err != nil {
			return err
		}
		nextNode, err := l.GetNode(index)
		if err != nil {
			return err
		}

		previousNode.next = &node
		nextNode.prev = &node

		node.next = nextNode
		node.prev = previousNode

		l.length++
	}
	return nil
}

func (l *DoublyLinkedList) Delete() error {
	// if linked list is empty return error
	if l.length == 0 {
		return ErrEmptyLinkedList
	}
	// set the last node as the previous node
	l.tail = l.tail.prev
	l.length--

	return nil
}

func (l *DoublyLinkedList) DeleteAt(index int) error {
	if index == 0 { // if the specified index is the first node
		l.head = l.head.next
	} else if index == l.length-1 { // if the specified index is the last node
		l.Delete()
	} else {
		previousNode, err := l.GetNode(index - 1)
		if err != nil {
			return err
		}
		nextNode, err := l.GetNode(index + 1)
		if err != nil {
			return err
		}
		previousNode.next = nextNode
		nextNode.prev = previousNode
	}

	l.length--
	return nil
}

func (l *DoublyLinkedList) GetNode(index int) (*Node, error) {
	// if linked list is empty or specified index is out of bounds
	if (index > l.length-1) || index < 0 {
		return nil, ErrOutOfBounds
	}

	ptr := l.head
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}

	return ptr, nil
}

// Display shows the nodes in the linked list
func (l *DoublyLinkedList) Display() {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d---><---", ptr.value)
		ptr = ptr.next
	}
	fmt.Print("\n")
}

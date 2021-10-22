package stack

const (
	ErrPopEmptyStack = StackError("cannot pop empty stack")
	ErrPushFullStack = StackError("cannot push into full stack")
)

type StackError string

func (s StackError) Error() string {
	return string(s)
}

type Stack struct {
	Top      int // is -1 if the stack is empty
	Elements []int
	Capacity int
}

// NewStack is a constructor to create a stack
func NewStack(capacity int, elements []int) *Stack {
	s := Stack{}
	if len(elements) != 0 {
		s.Elements = append(s.Elements, elements...)
	}
	s.Top = len(elements) - 1
	s.Capacity = capacity
	return &s
}

// Peek returns the index of the top element
// returns 0 if stack is empty (i.e Stack.Top = -1)
func (s Stack) Peek() int {
	if s.Top == -1 {
		return 0
	}
	return s.Elements[s.Top]
}

// Push adds one element to the stack
// returns error if stack is already full
func (s *Stack) Push(element int) error {
	if s.Top == s.Capacity-1 {
		return ErrPushFullStack
	}
	s.Elements = append(s.Elements, element)
	s.Top += 1
	return nil
}

// Pop removes an element from the stack
// returns error if the stack is empty
func (s *Stack) Pop() (int, error) {
	// attempt to pop empty stack returns error
	if s.Top == -1 {
		return 0, ErrPopEmptyStack
	}
	pop := s.Elements[s.Top]
	s.Elements = s.Elements[:s.Top] // remove last element by slicing
	s.Top -= 1
	return pop, nil
}

func (s Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s Stack) IsFull() bool {
	return s.Top == s.Capacity-1
}

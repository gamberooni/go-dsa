package stack

import "testing"

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack(5, []int{})
	got := stack.Peek()
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPush(t *testing.T) {
	t.Run("push into non-full stack", func(t *testing.T) {
		stack := NewStack(5, []int{})
		err := stack.Push(3)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		got := stack.Peek()
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("push into full stack", func(t *testing.T) {
		stack := NewStack(2, []int{4, 5})
		err := stack.Push(3)
		assertError(t, err, ErrPushFullStack)

		got := stack.Peek()
		want := 5

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("pop non-empty stack", func(t *testing.T) {
		stack := NewStack(5, []int{4, 5})
		got, err := stack.Pop()
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		want := 5

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("pop empty stack", func(t *testing.T) {
		stack := NewStack(5, []int{})
		got, err := stack.Pop()
		assertError(t, err, ErrPopEmptyStack)

		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestIsEmpty(t *testing.T) {
	t.Run("non-empty stack", func(t *testing.T) {
		stack := NewStack(5, []int{4, 5})
		got := stack.IsEmpty()
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack(5, []int{})
		got := stack.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("non-full stack", func(t *testing.T) {
		stack := NewStack(5, []int{4, 5})
		got := stack.IsFull()
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack(2, []int{4, 5})
		got := stack.IsFull()
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

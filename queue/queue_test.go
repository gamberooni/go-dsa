package queue

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
	q := NewQueue(5, []int{})
	got := q.Peek()
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestEnqueue(t *testing.T) {
	t.Run("non-full queue", func(t *testing.T) {
		q := NewQueue(5, []int{3})
		err := q.Enqueue(4)
		if err != nil {
			t.Fatal("didn't expect to get an error")
		}

		got := q.Elements[q.Rear]
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("full queue", func(t *testing.T) {
		q := NewQueue(2, []int{4, 5})
		err := q.Enqueue(4)
		assertError(t, err, ErrQueueIsFull)

		got := q.Elements[q.Rear]
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("non-empty queue", func(t *testing.T) {
		q := NewQueue(2, []int{4, 5})
		got, err := q.Dequeue()
		if err != nil {
			t.Fatal("didn't expect to get an error")
		}

		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue(2, []int{})
		got, err := q.Dequeue()
		assertError(t, err, ErrQueueIsEmpty)

		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("full queue", func(t *testing.T) {
		q := NewQueue(2, []int{4, 5})
		got := q.IsFull()
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("non-full queue", func(t *testing.T) {
		q := NewQueue(5, []int{4, 5})
		got := q.IsFull()
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue(5, []int{})
		got := q.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("non-empty queue", func(t *testing.T) {
		q := NewQueue(5, []int{4, 5})
		got := q.IsEmpty()
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

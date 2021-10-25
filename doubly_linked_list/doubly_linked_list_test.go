package doublylinkedlist

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

func TestInsert(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(4)
	l.Insert(5)

	node, err := l.GetNode(1)
	if err != nil {
		t.Fatal("didn't expect an error")
	}

	got := node.value
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestInsertAt(t *testing.T) {
	l := DoublyLinkedList{}
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)
	t.Run("valid index", func(t *testing.T) {
		l.InsertAt(1, 7)
		node, err := l.GetNode(1)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		got := node.value
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		err := l.InsertAt(5, 7)
		assertError(t, err, ErrOutOfBounds)
	})
}

func TestDelete(t *testing.T) {
	t.Run("non-empty linked list", func(t *testing.T) {
		l := DoublyLinkedList{}
		l.Insert(4)
		l.Insert(5)
		err := l.Delete()
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		node, err := l.GetNode(0)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		got := node.value
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("empty linked list", func(t *testing.T) {
		l := DoublyLinkedList{}
		err := l.Delete()
		assertError(t, err, ErrEmptyLinkedList)
	})
}

func TestDeleteAt(t *testing.T) {
	t.Run("valid index", func(t *testing.T) {
		l := DoublyLinkedList{}
		l.Insert(4)
		l.Insert(5)
		l.Insert(6)
		err := l.DeleteAt(1)
		if err != nil {
			t.Fatal("didn't expect an error")
		}
		node, err := l.GetNode(1)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		got := node.value
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		l := DoublyLinkedList{}
		l.Insert(4)
		l.Insert(5)
		l.Insert(6)
		err := l.DeleteAt(5)
		assertError(t, err, ErrOutOfBounds)
	})
}

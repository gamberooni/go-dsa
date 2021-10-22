package linkedlist

import (
	"reflect"
	"testing"
)

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
	l := LinkedList{}
	l.Insert(4)
	got := l.head.value
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestInsertAt(t *testing.T) {
	l := LinkedList{}
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	t.Run("valid index", func(t *testing.T) {
		err := l.InsertAt(1, 2)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		node, _ := l.GetNode(1)
		got := node.value
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		err := l.InsertAt(5, 2)
		assertError(t, err, ErrOutOfBounds)
	})
}

func TestDelete(t *testing.T) {
	l := LinkedList{}
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)
	l.Delete()
	node, _ := l.GetNode(1)
	got := node.value
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDeleteAt(t *testing.T) {
	l := LinkedList{}
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	t.Run("valid index", func(t *testing.T) {
		err := l.DeleteAt(0)
		if err != nil {
			t.Fatal("didn't expect to get an error")
		}

		node, _ := l.GetNode(0)
		got := node.value
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		err := l.DeleteAt(5)
		assertError(t, err, ErrOutOfBounds)
	})
}

func TestSearch(t *testing.T) {
	t.Run("search for a value present in the linked list", func(t *testing.T) {
		l := LinkedList{}
		l.Insert(6)
		l.Insert(5)
		l.Insert(4)

		got := l.Search(4)
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("search for a value not present in the linked list", func(t *testing.T) {
		l := LinkedList{}
		l.Insert(6)
		l.Insert(5)
		l.Insert(4)

		got := l.Search(7)
		want := -1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSort(t *testing.T) {
	l := LinkedList{}
	l.Insert(6)
	l.Insert(5)
	l.Insert(4)

	l.Sort()

	ptr := l.head
	got := []int{}
	for i := 0; i < l.length; i++ {
		got = append(got, ptr.value)
		ptr = ptr.next
	}

	want := []int{4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetNode(t *testing.T) {
	l := LinkedList{}
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	t.Run("given valid index", func(t *testing.T) {
		node, err := l.GetNode(1)
		if err != nil {
			t.Fatal("didn't expect an error")
		}

		got := node.value
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("given out of bounds index", func(t *testing.T) {
		_, err := l.GetNode(3)
		assertError(t, err, ErrOutOfBounds)
	})
}

package ll

import (
	"testing"
)

func TestGetNode(t *testing.T) {
	cases := []struct {
		pos   int
		value int
		want  int
	}{
		{pos: 0, value: 1, want: 1},
		{pos: 1, value: 2, want: 2},
		{pos: 2, value: 3, want: 3},
		{pos: 3, value: 4, want: 4},
	}

	l := LinkedList{}

	for _, c := range cases {
		l.Insert(c.value)
		node := l.GetNode(c.pos)
		if node.value != c.want {
			t.Errorf("GetNode error, value : %d should be %d", node.value, c.want)
		}
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		value int
		want  int
	}{
		{value: 1, want: 1},
		{value: 2, want: 2},
		{value: 3, want: 3},
		{value: 4, want: 4},
	}

	l := LinkedList{}
	for i, c := range cases {
		l.Insert(c.value)
		if l.len != i+1 {
			t.Errorf("Insert error, length : %d should be %d", l.len, i+1)
		}

		if l.GetNode(i).value != c.want {
			t.Errorf("Insert error, value : %d should be %d", l.head.value, c.want)
		}
	}
}

func TestInsertAt(t *testing.T) {
	cases := []struct {
		pos   int
		value int
		want  int
	}{
		{pos: 0, value: 0, want: 0},
		{pos: 1, value: 2, want: 2},
		{pos: 2, value: 3, want: 3},
		{pos: 3, value: 4, want: 4},
		{pos: 7, value: 5, want: 5},
	}

	l := LinkedList{}

	l.Insert(11)
	l.Insert(22)
	l.Insert(33)
	l.Insert(44)

	for _, c := range cases {
		l.InsertAt(c.pos, c.value)
		node := l.GetNode(c.pos)
		if node.value != c.want {
			t.Errorf("InsertAt error, value : %d should be %d", node.value, c.want)
		}
	}
	//
	//l.PrintL()
	//fmt.Printf("%+v\n", l)
}

func TestSearch(t *testing.T) {
	cases := []struct {
		value int
		want  int
	}{
		{value: 1, want: 0},
		{value: 2, want: 1},
		{value: 3, want: 2},
		{value: 4, want: 3},
	}

	l := LinkedList{}
	for _, i := range cases {
		l.Insert(i.value)
	}

	for _, c := range cases {
		nodePos := l.Search(c.value)
		if nodePos != c.want {
			t.Errorf("Search error, want: %d, got: %d\n", c.want, nodePos)
		}
	}

	nodePos := l.Search(100)
	if nodePos != -1 {
		t.Errorf("Search error, want: %d, got: %d\n", -1, nodePos)
	}
}

func TestDeleteVal(t *testing.T) {
	l := LinkedList{}

	err := l.DeleteVal(222)
	if err == nil {
		t.Errorf("DeleteVal error: Got nil but want an error")
	}

	l.Insert(11)
	l.Insert(22)
	l.Insert(33)
	l.Insert(44)

	// cases
	// - delete value that exists
	// - delete value that not exists

	cases := []struct {
		value int
		//pos   int
		//want int
	}{

		{value: 11},
		{value: 33},
		{value: 44},
	}

	err = l.DeleteVal(22)
	if err != nil {
		t.Fatalf("DeleteVal error: We got not expected error")
	}

	for i, c := range cases {
		if l.GetNode(i).value != c.value {
			t.Errorf("DeleteVal error: got %d, want: %d\n",
				l.GetNode(i).value, c.value)
		}
	}

	err = l.DeleteVal(222)
	if err == nil {
		t.Errorf("DeleteVal error: Got nil but want an error")
	}

}

func TestDeleteAt(t *testing.T) {
	cases := []struct {
		value int
		pos   int
		want  int
	}{
		{value: 1, pos: 0, want: 2},
		{value: 2, pos: 0, want: 3},
		{value: 3, pos: 0, want: 4},
		{value: 4, pos: 0, want: 0},
	}

	l := LinkedList{}
	for _, i := range cases {
		l.Insert(i.value)
	}
	//l.PrintL()

	origLen := l.len

	for _, c := range cases {
		err := l.DeleteAt(c.pos)
		if err != nil {
			t.Errorf("Got unexpected error: %s\n", err)
		}

		if l.len != origLen-1 {
			t.Errorf("DeleteAt error, len: %d, want: %d\n", l.len, origLen)
		}

		if l.head.value != c.want {
			t.Errorf("DeleteAt error, want value: %d, got: %d\n",
				c.want, l.head.value)
		}

		origLen--
	}

	l = LinkedList{}

	l.Insert(11)
	l.Insert(22)
	l.Insert(33)
	l.Insert(44)

	err := l.DeleteAt(1)
	if err != nil {
		t.Errorf("Got unexpected error: %s\n", err)
	}

	if l.len != 3 {
		t.Errorf("DeleteAt error, len: %d, want: %d\n", l.len, 3)
	}

	if l.head.value != 11 {
		t.Errorf("DeleteAt error, want value: %d, got: %d\n",
			11, l.head.value)
	}

	if l.head.next.value != 33 {
		t.Errorf("DeleteAt error, want value: %d, got: %d\n",
			33, l.head.value)
	}
}

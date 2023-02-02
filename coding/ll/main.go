package ll

import (
	"errors"
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	current := l.head
	for i := 0; i < l.len; i++ {

		if current.next == nil {
			current.next = &n
			l.len++
			break
		} else {
			current = current.next
		}

	}
}

func (l LinkedList) GetNode(pos int) *Node {

	if pos == 0 {
		return l.head
	}

	//fmt.Printf("pos is: %d\n", pos)
	//fmt.Printf("l.len: %d\n", l.len)
	//fmt.Printf("l.head: %+v\n", l.head)

	for i := 0; ; i++ {
		//fmt.Printf("i: %d\n", i)
		if i == pos {
			return l.head
		} else {
			l.head = l.head.next
		}
	}
}

func (l *LinkedList) InsertAt(pos int, value int) {
	if l.len < pos {
		return
	}

	n := Node{}
	n.value = value

	if pos == 0 {
		l.head, n.next = &n, l.head
		l.len++
		return
	}

	if pos > l.len {
		return
	}

	node := l.GetNode(pos)
	//fmt.Printf("node value: %d\n", node.value)

	prevNode := l.GetNode(pos - 1)
	//fmt.Printf("prev node value: %d\n", prevNode.value)

	prevNode.next = &n
	n.next = node
	l.len++

}

func (l LinkedList) Search(val int) int {
	// return node position by val

	current := l.head
	for i := 0; i < l.len; i++ {
		if current.value == val {
			return i
		} else {
			current = current.next
		}
	}
	return -1
}

func (l *LinkedList) DeleteAt(pos int) error {

	if pos < 0 {
		return errors.New("position can't be lower zero")
	}

	if pos >= l.len {
		return errors.New("position can't be higher len of ll")
	}

	if l.len == 0 {
		return errors.New("ll is empty")
	}

	if l.len == 1 {
		l.head = &Node{}
		l.len--
		return nil
	}

	//fmt.Println("===============")
	//l.PrintL()
	//fmt.Printf("len is: %d\n", l.len)

	if pos == 0 {
		l.head = l.head.next
	} else {
		prevNode := l.GetNode(pos - 1)
		nextNode := l.GetNode(pos + 1)
		prevNode.next = nextNode
	}

	l.len--
	return nil
}

func (l LinkedList) PrintL() {

	if l.len == 0 {
		return
	}
	current := l.head

	for i := 0; i < l.len; i++ {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	//d := Node{4, nil}
	//c := Node{3, &d}
	//b := Node{2, &c}
	//a := Node{1, &b}
	//
	//current := a
	//for current.next != nil {
	//	fmt.Println(current.value)
	//	current = *current.next
	//}

	//fmt.Println("Use LinkList structure")

	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)

	//for i := 0; i < ll.len; i++ {
	//	fmt.Println(ll.head.value)
	//	ll.head = ll.head.next
	//}

	ll.PrintL()

	ll.InsertAt(1, 11)

	ll.PrintL()

	fmt.Printf("Inset 0 at 0 pos\n")
	ll.InsertAt(0, 0)

	ll.PrintL()

}

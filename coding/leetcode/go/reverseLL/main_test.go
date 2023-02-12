package reverseLL

import (
	"fmt"
	"testing"
)

var test = []struct {
	got  []int
	want []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
	},
	{
		[]int{1, 2},
		[]int{2, 1},
	},
	{
		[]int{},
		[]int{},
	},
}

func (l *ListNode) Insert(val int) {
	n := &ListNode{val, nil}
	if l == nil {
		l.Val = val
		l.Next = nil
	} else {
		l.Next = n
	}
}

func compareLL(one *ListNode, two *ListNode) bool {

	//fmt.Printf("one: %v, two: %v\n", one, two)

	for one.Next != nil && two.Next != nil {
		if one.Val != two.Val {
			return false
		}

		one = one.Next
		two = two.Next
	}
	return true

}

func returnSliceFromLL(ll *ListNode) []int {
	var sliceLL []int

	if ll == nil {
		return sliceLL
	}

	for ll.Next != nil {
		sliceLL = append(sliceLL, ll.Val)
		ll = ll.Next
	}
	return sliceLL
}

func PrintL(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func TestReverseList(t *testing.T) {

	ll := &ListNode{}
	reversedLL := &ListNode{}

	for _, n := range test {
		for _, v := range n.got {
			ll.Insert(v)
		}

		PrintL(ll)

		head := reverseList(ll)

		for _, v := range n.want {
			reversedLL.Insert(v)
		}

		if !compareLL(ll, reversedLL) {
			t.Errorf("TestReverseList: Want %v, got: %v\n",
				n.want, returnSliceFromLL(head))
		}
		ll = &ListNode{}
		reversedLL = &ListNode{}
	}
}

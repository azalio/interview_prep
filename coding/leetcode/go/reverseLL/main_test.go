package reverseLL

import (
	"testing"
)

var test = []struct {
	got  []int
	want []int
}{
	//{
	//	[]int{1, 2, 3, 4, 5},
	//	[]int{5, 4, 3, 2, 1},
	//},
	//{
	//	[]int{1, 2},
	//	[]int{2, 1},
	//},
	{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
	},
	//{
	//	[]int{},
	//	[]int{},
	//},
}

func TestReverseList(t *testing.T) {

	//e := &ListNode{5, nil}
	//d := &ListNode{4, e}
	c := &ListNode{3, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	//
	//PrintL(a)
	//fmt.Println("==================")
	reverseLL := reverseList(a)

	t.Errorf("%v\n", reverseLL)

	//PrintL(reverseLL)

}

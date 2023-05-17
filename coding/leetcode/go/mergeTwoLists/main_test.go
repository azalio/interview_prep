package mergeTwoLists

import (
	"fmt"
	"testing"
)

var test = []struct {
	list1 []int
	list2 []int
	want  []int
}{
	{
		[]int{1, 2, 4},
		[]int{1, 3, 4},
		[]int{1, 1, 2, 3, 4, 4},
	},
	//{
	//	[]int{},
	//	[]int{},
	//	[]int{},
	//},
	//{
	//	[]int{},
	//	[]int{0},
	//	[]int{0},
	//},
}

func createLL(s []int) ListNode {
	if len(s) == 0 {
		return ListNode{}
	}

	if len(s) == 1 {
		return ListNode{s[0], nil}
	}

	ll := ListNode{s[len(s)-1], nil}
	//fmt.Printf("%v\n", ll)

	for i := len(s) - 2; i >= 0; i-- {
		oldLL := ll
		ll = ListNode{s[i], &oldLL}
		//fmt.Printf("%v\n", ll)

	}

	return ll
}

func printLL(ll ListNode) {
	for {
		fmt.Println(ll.Val)
		if ll.Next == nil {
			return
		}
		ll = *ll.Next
	}
}

func TestMergeTwoLists(t *testing.T) {

	for _, tt := range test {
		ll1 := createLL(tt.list1)
		ll2 := createLL(tt.list2)

		//fmt.Println(ll1.Next.Next.Val)
		//printLL(ll1)
		//printLL(ll2)

		result := mergeTwoLists(&ll1, &ll2)
		printLL(*result)
	}
}

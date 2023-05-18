package mergeTwoLists

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	a := list1
	b := list2

	head := ListNode{}
	curr := &head

	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}
		curr = curr.Next
	}

	if a != nil {
		curr.Next = a
	}

	if b != nil {
		curr.Next = b
	}

	return head.Next
}

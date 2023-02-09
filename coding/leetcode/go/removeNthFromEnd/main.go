package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a linked list,
// remove the nth node from the end of the list and return its head.

func printLL(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lead := &ListNode{0, head}
	fast := lead
	slow := lead

	printLL(lead)

	for ; n > 0; n-- {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return lead.Next
}

func main() {
	e := &ListNode{5, nil}
	d := &ListNode{4, e}
	c := &ListNode{3, d}
	b := &ListNode{2, c}
	a := &ListNode{1, b}

	_ = removeNthFromEnd(a, 2)
	//
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}

}

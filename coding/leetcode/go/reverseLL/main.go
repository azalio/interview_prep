package reverseLL

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintL(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func reverseList(head *ListNode) *ListNode {

	//PrintL(head)

	current := head
	var next *ListNode = nil
	var prev *ListNode = nil

	for current != nil {
		fmt.Println("+++++++++")

		fmt.Printf("current: %v\n", current)
		next = current.Next
		fmt.Printf("next: %v\n", next)
		current.Next = prev
		fmt.Printf("current.next: %v\n", current.Next)
		prev = current
		fmt.Printf("prev: %v\n", prev)
		current = next
		fmt.Printf("current: %v\n", current)
	}

	return prev

}

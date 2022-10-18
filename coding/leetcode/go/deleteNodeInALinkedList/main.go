package deleteNodeInALinkedList

// https://leetcode.com/problems/delete-node-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	node.Next = node.Next.Next

}

func deleteNode2(node *ListNode) {
	*node = *node.Next
}

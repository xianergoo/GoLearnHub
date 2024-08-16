package swapnodeinpairs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	prev := head
	next := head.Next.Next

	cur.Next = prev
	prev.Next = swapPairs(next)
	return cur
}

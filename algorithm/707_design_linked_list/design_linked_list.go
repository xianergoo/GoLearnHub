package design_linked_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (m *MyLinkedList) Print() {
	cur := m.head
	l := []int{}
	for cur != nil {
		l = append(l, cur.Val)
		cur = cur.Next
	}
	fmt.Println(m.size, l)
}

func (m *MyLinkedList) Get(index int) int {
	if m == nil || m.size <= 0 || index < 0 {
		return -1
	}
	cur := m.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur != nil {
		return cur.Val
	}
	return -1
}

func (m *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{Val: val}
	newHead.Next = m.head
	m.head = newHead
	m.size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	newHead := &ListNode{Val: val}
	if m.size == 0 {
		m.head = newHead
	} else {
		cur := m.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newHead
	}
	m.size++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		m.AddAtHead(val)
		return
	}
	if index > m.size {
		return
	}

	cur := m.head
	for i := 0; i < index-1; i++ {
		if cur == nil {
			break
		}
		cur = cur.Next
	}
	if cur != nil {
		newHead := &ListNode{Val: val}
		newHead.Next = cur.Next
		cur.Next = newHead
		m.size++
	}

}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= m.size {
		return
	}
	if index == 0 {
		m.head = m.head.Next
	} else {
		current := m.head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	m.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	list := Constructor()
	list.Print()

	list.AddAtHead(2)
	list.Print()

	// list.AddAtTail(3)
	// list.Print()

	list.AddAtIndex(0, 1)
	list.Print()

	a := list.Get(1)
	fmt.Println(a)

	// list.DeleteAtIndex(1)
	// list.Print()

	// b := list.Get(1)
	// fmt.Println(b)
}

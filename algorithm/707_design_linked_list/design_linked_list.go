package main

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

func (this *MyLinkedList) Print() {
	cur := this.head
	l := []int{}
	for cur != nil {
		l = append(l, cur.Val)
		cur = cur.Next
	}
	fmt.Println(this.size, l)
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.size <= 0 || index < 0 {
		return -1
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur != nil {
		return cur.Val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{Val: val}
	newHead.Next = this.head
	this.head = newHead
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newHead := &ListNode{Val: val}
	if this.size == 0 {
		this.head = newHead
	} else {
		cur := this.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newHead
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.size {
		return
	}

	cur := this.head
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
		this.size++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		this.head = this.head.Next
	} else {
		current := this.head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	this.size--
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

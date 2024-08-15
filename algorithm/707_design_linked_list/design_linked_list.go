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
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Print() {
	cur := this.head
	l := []int{}
	for i := 0; i < this.size; i++ {
		l = append(l, cur.Val)
		cur = cur.Next
	}
	fmt.Println(l)
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index > this.size-1 || index < 0 {
		return -1
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{Val: val}
	newHead.Next = this.head
	this.head = newHead
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	// this.AddAtIndex(this.size, val)
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
	if index < 0 || index > this.size {
		return
	}

	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	newHead := &ListNode{Val: val}
	newHead.Next = pre.Next
	pre.Next = newHead
	this.size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.size--
}

func main() {
	list := Constructor()
	list.Print()

	list.AddAtHead(1)
	list.Print()

	list.AddAtTail(3)
	list.Print()

	// list.AddAtIndex(1, 2)
	// list.Print()

	a := list.Get(1)
	fmt.Println(a)

	// list.DeleteAtIndex(1)
	// list.Print()

	b := list.Get(1)
	fmt.Println(b)
}

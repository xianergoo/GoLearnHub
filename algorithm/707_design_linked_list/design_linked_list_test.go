package design_linked_list

import (
	"testing"
)

func TestMyLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.index); got != tt.want {
				t.Errorf("MyLinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtHead(tt.args.val)
		})
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtTail(tt.args.val)
		})
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.DeleteAtIndex(tt.args.index)
		})
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtIndex(tt.args.index, tt.args.val)
		})
	}
}

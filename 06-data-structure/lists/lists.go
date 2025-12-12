package lists

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Len  uint
}

func CreateList[T any](data T) *List[T] {
	head := &Node[T]{data, nil}
	return &List[T]{head, head, 1}
}

func CreateNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil}
}

func RPush[T any](list *List[T], data T) *List[T] {
	node := &Node[T]{data, nil}
	list.Tail.Next = node
	list.Tail = node
	list.Len++

	return list
}

func LPush[T any](list *List[T], data T) *List[T] {
	node := &Node[T]{data, list.Head}
	list.Head = node
	list.Len++

	return list
}

func LLen[T any](list *List[T]) uint {
	return list.Len
}

func LPrint[T any](list *List[T]) {
	head := list.Head
	for head != nil {
		fmt.Printf("val: %v\n", head.Val)
		head = head.Next
	}
}

package main

import (
	"fmt"
	"learning/errhandling/lists"
	"learning/errhandling/tree"
)

func main() {

	list := lists.CreateList(10)
	lists.RPush(list, 11)
	lists.RPush(list, 12)
	lists.RPush(list, 13)
	lists.RPush(list, 14)
	lists.LPush(list, -1)

	head := list.Head
	tail := list.Tail
	length := list.Len

	fmt.Printf("head data: %d\n", head.Val)
	fmt.Printf("tail data: %d\n", tail.Val)
	fmt.Printf("list len: %d\n", length)

	lists.LPrint(list)

	bt := tree.CreateEmptyTree()
	bt = tree.Add(bt, 10)
	bt = tree.Add(bt, 50)
	bt = tree.Add(bt, 5)
	bt = tree.Add(bt, 20)
	bt = tree.Add(bt, -1)
	bt = tree.Add(bt, 2)
	bt = tree.Add(bt, 3)
	bt = tree.Add(bt, 99)

	tree.InOrdem(bt)
}

package main

import (
	"errors"
	"fmt"
)

type TreeNode struct { // Class
	left  *TreeNode
	val   int
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == value {
		return errors.New("This node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func traverse(t *TreeNode) {
	if t == nil {
		return
	}
	traverse(t.left)
	fmt.Print(t.val, " ") // 1 - A
	traverse(t.right)
}

func main() {

	head := &TreeNode{left: nil, val: 1, right: nil} // из-за того что исп. new первый элемент 0

	a := head.Insert(1)

	if a != nil {
		fmt.Print(a.Error(), " ")
	}

	head.Insert(20)
	head.Insert(13)
	head.Insert(41)
	head.Insert(25)
	head.Insert(36)
	head.Insert(78)
	head.Insert(20)
	head.Insert(13)
	head.Insert(113)

	traverse(head)

	//fmt.Println(head.left)

}

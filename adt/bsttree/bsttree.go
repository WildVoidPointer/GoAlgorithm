package bsttree

import (
	"cmp"
	"errors"
	"fmt"
)

type BinarySearchTreeNode[T cmp.Ordered] struct {
	data  T
	left  *BinarySearchTreeNode[T]
	right *BinarySearchTreeNode[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	root *BinarySearchTreeNode[T]
}

func NewBSTNode[T cmp.Ordered](v T) *BinarySearchTreeNode[T] {
	return &BinarySearchTreeNode[T]{
		data:  v,
		left:  nil,
		right: nil,
	}
}

func NewBST[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) InsertByRecursion(d T) bool {
	var err error
	bst.root, err = bst.insertByRecursionHelper(bst.root, d)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (bst *BinarySearchTree[T]) insertByRecursionHelper(
	node *BinarySearchTreeNode[T],
	d T,
) (*BinarySearchTreeNode[T], error) {
	if node == nil {
		return NewBSTNode(d), nil
	}

	if d < node.data {
		node.left, _ = bst.insertByRecursionHelper(node.left, d)
	} else if d > node.data {
		node.right, _ = bst.insertByRecursionHelper(node.right, d)
	} else {
		return node, errors.New("BinarySearchTree: " +
			"This value already exists in the BinarySearchTree")
	}
	return node, nil
}

func (bst *BinarySearchTree[T]) InOrderPrintln() {
	fmt.Print("BinarySearchTree: {  ")
	bst.inOrderPrintlnHelper(bst.root)
	fmt.Print("}\n")
}

func (bst *BinarySearchTree[T]) inOrderPrintlnHelper(
	node *BinarySearchTreeNode[T],
) {
	if node != nil {
		bst.inOrderPrintlnHelper(node.left)
		fmt.Printf("%v  ", node.data)
		bst.inOrderPrintlnHelper(node.right)
	}
}

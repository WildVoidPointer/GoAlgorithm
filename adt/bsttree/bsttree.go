package bsttree

import (
	"cmp"
	"errors"
	"fmt"
)

type BSTreeNode[T cmp.Ordered] struct {
	data  T
	left  *BSTreeNode[T]
	right *BSTreeNode[T]
}

type BSTree[T cmp.Ordered] struct {
	root *BSTreeNode[T]
}

func NewBSTNode[T cmp.Ordered](v T) *BSTreeNode[T] {
	return &BSTreeNode[T]{
		data:  v,
		left:  nil,
		right: nil,
	}
}

func NewBST[T cmp.Ordered]() *BSTree[T] {
	return &BSTree[T]{}
}

func (bst *BSTree[T]) Search(d T) (BSTreeNode[T], bool) {

	if bst.root == nil {
		var v BSTreeNode[T]
		return v, false
	}

	curr := bst.root

	for curr != nil {
		if d > curr.data {
			curr = curr.right
		} else if d < curr.data {
			curr = curr.left
		} else {
			return *curr, true
		}
	}

	var v BSTreeNode[T]
	return v, false
}

func (bst *BSTree[T]) InsertByIter(d T) bool {
	if bst.root == nil {
		bst.root = NewBSTNode(d)
		return true
	}

	curr := bst.root
	for curr != nil {
		if d > curr.data {
			if curr.right == nil {
				curr.right = NewBSTNode(d)
				return true
			}
			curr = curr.right
		} else if d < curr.data {
			if curr.left == nil {
				curr.left = NewBSTNode(d)
				return true
			}
			curr = curr.left
		} else {
			return false
		}
	}

	return true
}

func (bst *BSTree[T]) InsertByRecur(d T) bool {
	var err error
	bst.root, err = bst.insertByRecurHelper(bst.root, d)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (bst *BSTree[T]) Remove(d T) (BSTreeNode[T], bool) {
	if bst == nil {
		var v BSTreeNode[T]
		return v, false
	}

	var prev *BSTreeNode[T]

	curr := bst.root

	for curr != nil && curr.data == d {
		prev = curr

		if d > curr.data {
			curr = curr.right
		} else if d < curr.data {
			curr = curr.left
		}
	}
}

func (bst *BSTree[T]) InOrderPrintln() {
	fmt.Print("BinarySearchTree: {  ")
	bst.inOrderPrintlnHelper(bst.root)
	fmt.Print("}\n")
}

func (bst *BSTree[T]) inOrderPrintlnHelper(
	node *BSTreeNode[T],
) {
	if node != nil {
		bst.inOrderPrintlnHelper(node.left)
		fmt.Printf("%v  ", node.data)
		bst.inOrderPrintlnHelper(node.right)
	}
}

func (bst *BSTree[T]) insertByRecurHelper(
	node *BSTreeNode[T],
	d T,
) (*BSTreeNode[T], error) {

	if node == nil {
		return NewBSTNode(d), nil
	}

	if d < node.data {
		node.left, _ = bst.insertByRecurHelper(node.left, d)
	} else if d > node.data {
		node.right, _ = bst.insertByRecurHelper(node.right, d)
	} else {
		return node, errors.New("BSTree: " +
			"This value already exists in the BinarySearchTree")
	}
	return node, nil
}

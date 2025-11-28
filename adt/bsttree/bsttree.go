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

func (bst *BSTree[T]) RemoveByIter(d T) (BSTreeNode[T], bool) {

	if bst == nil {
		var v BSTreeNode[T]
		return v, false
	}

	// 存储被删除节点前驱
	var prev *BSTreeNode[T]

	// 存储被删除节点以供返回
	var toDelete BSTreeNode[T]

	// 进行节点查找
	curr := bst.root

	for curr != nil && curr.data != d {
		prev = curr

		if d > curr.data {
			curr = curr.right
		} else if d < curr.data {
			curr = curr.left
		}
	}

	// 树中无此节点直接返回
	if curr == nil {
		var v BSTreeNode[T]
		return v, false
	} else {
		toDelete = *curr
	}

	// 如果同时具有左右子树
	if curr.left != nil && curr.right != nil {

		inorderFirst := curr
		var inorderFirstPrev *BSTreeNode[T]

		// 寻找中序第一个节点
		for inorderFirst.left != nil {
			inorderFirstPrev = inorderFirst
			inorderFirst = inorderFirst.left
		}

		// 将中序第一节点数据设置给被删除节点
		curr.setFromNode(inorderFirst)

		// 变换删除节点为中序第一节点
		curr = inorderFirst
		prev = inorderFirstPrev
	}

	var child *BSTreeNode[T]

	if curr.left != nil {
		child = curr.left
	} else {
		child = curr.right
	}

	if prev == nil {
		bst.root = child
	} else {
		if prev.left == curr {
			prev.left = child
		} else {
			prev.right = child
		}
	}

	return toDelete, true
}


func (node *BSTreeNode[T]) setFromNode(n *BSTreeNode[T]) {
	node.data = n.data
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

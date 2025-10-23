package avltree

import "cmp"

type AvlTreeNode[T cmp.Ordered] struct {
	data 	T
	left 	*AvlTreeNode[T]
	right 	*AvlTreeNode[T]
}

type AvlTree[T cmp.Ordered] struct {
	root 	*AvlTreeNode[T]
}

func NewAvlTreeNode[T cmp.Ordered](d T) *AvlTreeNode[T] {
	return &AvlTreeNode[T]{
		data: 	d,
		left: 	nil,
		right: 	nil,
	}
}


func NewAvlTree[T cmp.Ordered]() *AvlTree[T] {
	return &AvlTree[T]{}
} 

package bsttree

func RunBinarySearchTreeTest() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	bst := NewBST[int]()

	for _, v := range arr {
		bst.InsertByRecursion(v)
	}

	bst.InOrderPrintln()
}

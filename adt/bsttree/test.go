package bsttree

func RunBSTreeTest() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	bst_recur := NewBST[int]()

	for _, v := range arr {
		bst_recur.InsertByRecur(v)
	}

	bst_recur.InOrderPrintln()

	bst_iter := NewBST[int]()

	for _, v := range arr {
		bst_iter.InsertByIter(v)
	}

	bst_iter.InOrderPrintln()

	bst_iter.RemoveByIter(5)

	bst_iter.InOrderPrintln()
}

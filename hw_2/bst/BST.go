package main

type BST struct {
	data  int
	left  *BST
	right *BST
}

func (bst *BST) Add(data int) {
	root := bst
	for {
		if data == root.data {
			return
		}
		if data < root.data {
			if root.left == nil {
				root.left = &BST{data: data, left: nil, right: nil}
				break
			}
			root = root.left
		}
		if data > root.data && root.right != nil {
			if root.right == nil {
				root.right = &BST{data: data, left: nil, right: nil}
				break
			}
			root = root.right
		}
	}
}

func (bst *BST) isExist(data int) bool {
	root := bst
	for {
		if data == root.data {
			return true
		}
		if data < root.data {
			if root.left == nil {
				return false
			}
			root = root.left
			continue
		}
		if data > root.data {
			if root.right == nil {
				return false
			}
			root = root.right
			continue
		}
	}
}

func (bst *BST) Delete(data int) {
	var root, cur, parent *BST
	root = bst
	cur = root

	for cur != nil {
		if data < cur.data {
			parent = cur
			cur = cur.left
			continue
		} else if data > cur.data {
			parent = cur
			cur = cur.right
			continue
		} else {
			break
		}
	}

	if cur == nil {
		return
	}

	if cur.right == nil {
		if parent != nil {
			if cur == parent.left {
				parent.left = cur.left
			} else {
				parent.right = cur.left
			}
		} else {
			if cur.left != nil {
				root.data = cur.left.data
				root.left = cur.left.left
			} else {
				root = cur.left
			}
		}
	} else {
		parent = cur
		min := cur.right

		for min.left != nil {
			parent = min
			min = min.left
		}

		cur.data = min.data
		if min == parent.left {
			parent.left = min.right
		} else {
			parent.right = min.right
		}
	}
}

func NewBST(data int) *BST {
	return &BST{data: data, left: nil, right: nil}
}

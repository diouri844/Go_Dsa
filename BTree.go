package main

import "fmt"

// crete A node type :

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// insert new node :

func (n *TreeNode) insert(item int) {
	if item > n.data {
		if n.left == nil {
			n.left = &TreeNode{data: item}
		} else {
			n.left.insert(item)
		}
	} else if item < n.data {
		if n.right == nil {
			n.right = &TreeNode{data: item}
		} else {
			n.right.insert(item)
		}
	}
}

// display method :
func (n *TreeNode) ShowTree() {
	if n != nil {
		fmt.Println(n.data)
	}
	if n.left != nil {
		fmt.Println(" left --> ", n.left.data)
		n.left.ShowTree()
	}
	if n.right != nil {
		fmt.Println(" right ---> ", n.right.data)
		n.right.ShowTree()
	}
}

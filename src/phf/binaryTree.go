package phf

import (
	//"fmt"
)

// Define the Node structure
type Node struct {
	key   uint64
	left  *Node
	right *Node
}

// Define the BinaryTree structure
type BinaryTree struct {
	root *Node
}

// Insert a key into the binary tree
func (bt *BinaryTree) Insert(key uint64) {
	if bt.root == nil {
		bt.root = &Node{key: key}
	} else {
		bt.root.insert(key)
	}
}

// Insert a key into the node
func (n *Node) insert(key uint64) {
	if key < n.key {
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.insert(key)
		}
	} else if key > n.key {
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.insert(key)
		}
	}
}

// Search for a key in the binary tree
func (bt *BinaryTree) Search(key uint64) bool {
	return bt.root != nil && bt.root.search(key)
}

// Search for a key in the node
func (n *Node) search(key uint64) bool {
	if key == n.key {
		return true
	}
	if key < n.key {
		if n.left == nil {
			return false
		}
		return n.left.search(key)
	}
	if n.right == nil {
		return false
	}
	return n.right.search(key)
}

// InOrder traversal of the binary tree
func (bt *BinaryTree) InOrder() {
	if bt.root != nil {
		bt.root.inOrder()
	}
}

// InOrder traversal of the node
func (n *Node) inOrder() {
	if n.left != nil {
		n.left.inOrder()
	}
	//fmt.Println(n.key)
	if n.right != nil {
		n.right.inOrder()
	}
}

/*
func main() {
	// Create a new BinaryTree
	bt := &BinaryTree{}

	// Example keys to insert
	keys := []uint64{1000, 500, 1500, 250, 750, 1250, 1750}

	// Insert keys into the binary tree
	for _, key := range keys {
		bt.Insert(key)
	}

	// Search for keys in the binary tree
	searchKeys := []uint64{500, 1250, 2000}
	for _, key := range searchKeys {
		found := bt.Search(key)
		if found {
			fmt.Printf("Key %d found in the binary tree.\n", key)
		} else {
			fmt.Printf("Key %d not found in the binary tree.\n", key)
		}
	}

	// InOrder traversal of the binary tree
	fmt.Println("InOrder traversal of the binary tree:")
	bt.InOrder()
}
*/

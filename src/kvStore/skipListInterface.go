package keyValueStore

import (
	"fmt"
	"math/rand"
)

type DataSkipList interface {
	NewSkipList()
	Insert()
	PrintList()
}

type SkipList struct {
	head   *Node
	height int
}

type Node struct {
	NextNode []*Node
	Value    int
	key      uint64
}

const maxHeight = 5

func NewSkipList() *SkipList {
	head := NewNode(0, 1, maxHeight)
	return &SkipList{
		head:   head,
		height: 1,
	}
}

func NewNode(key uint64, value int, height int) *Node {
	return &Node{
		NextNode: make([]*Node, height),
		Value:    value,
		key:      key,
	}
}

func (skipList *SkipList) Insert(key uint64, node Node) {
	updatePointer := make([]*Node, maxHeight)
	currentNode := skipList.head

	for i := skipList.height; i >= 0; i-- {
		for currentNode.NextNode[i] != nil && currentNode.NextNode[i].key < key {
			currentNode = currentNode.NextNode[i]
		}
		updatePointer[i] = currentNode
	}

	height := rand.Intn(maxHeight)

	if height > skipList.height {
		// Extend updatePointer to accommodate the new height
		newUpdatePointer := make([]*Node, height+1)
		copy(newUpdatePointer, updatePointer)
		for i := skipList.height + 1; i <= height; i++ {
			newUpdatePointer[i] = skipList.head
		}
		updatePointer = newUpdatePointer
		skipList.height = height
	}

	newNode := NewNode(key, node.Value, height)

	for i := 0; i < height; i++ {
		newNode.NextNode[i] = updatePointer[i].NextNode[i]
		updatePointer[i].NextNode[i] = newNode
	}
}

func (skipList *SkipList) PrintList() {
	for height := skipList.height; height >= 0; height-- {
		current := skipList.head.NextNode[height]
		fmt.Printf("Level %d: ", height)
		for current != nil {
			fmt.Printf("%d ", current.key)
			current = current.NextNode[height]
		}
		fmt.Println()
	}
}

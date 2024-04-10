package main

import (
	keyValueStore "github.com/RaphaelCarl1/dmds/src/kvStore"
)

func main() {
	// Initialize SkipList
	skipList := &keyValueStore.SkipList{
		Head:   &keyValueStore.Node{NextNode: make([]*keyValueStore.Node, keyValueStore.MaxHeight)},
		Height: 1,
		Length: 0,
	}

	// Example usage: Insert key-value pairs into SkipList
	skipList.Put(1, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	skipList.Put(3, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	skipList.Put(2, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	skipList.Put(5, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	skipList.Put(4, [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
}

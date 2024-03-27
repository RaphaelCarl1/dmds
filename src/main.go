package main

import (
	keyValueStore "github.com/RaphaelCarl1/dmds/src/kvStore"
)

func main() {

	kvStore := keyValueStore.Create(5)
	kvStore.Put(3, [10]byte{'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '3', '4'})
	kvStore.Put(2, [10]byte{'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't'})
	kvStore.Put(1, [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'})
	kvStore.Put(5, [10]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'})
	kvStore.Put(4, [10]byte{'0', '9', '8', '7', '6', '5', '4', '3', '2', '1'})

	skipList := keyValueStore.NewSkipList()
	skipList.Insert(1, keyValueStore.Node{Value: 10})
	skipList.Insert(5, keyValueStore.Node{Value: 30})
	skipList.Insert(3, keyValueStore.Node{Value: 20})
	skipList.PrintList()
}

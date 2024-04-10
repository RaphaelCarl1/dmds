package keyValueStore

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"math/rand"
)

type Data interface {
	Put(uint64, [10]byte)
	Get(uint64) uint64
	CreateSSTable() //split it up for data manipulation and writing it to a file
}

type Control interface {
	Create(uint64)
	Open(string, string)
	Close()
	Delete(string)
}

type SSTable struct {
	minKey       uint64
	maxKey       uint64
	shallowIndex []uint64
	//skipList     map[uint64][10]byte
}

type SkipList struct {
	Head   *Node
	Height int
	Length int
}

type Node struct {
	NextNode []*Node
	Value    [10]byte
	Key      uint64
}

//defines the height of the SkipList
var MaxHeight = 5
//defines the length of a SkipList
var MaxLength = 5

//creates a new, empty SkipList
func Create(maxHeight int) *SkipList {
	MaxHeight = maxHeight
	head := NewNode(0, [10]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}, maxHeight)
	return &SkipList{
		Head:   head,
		Height: 1,
	}
}

//creates a new Node for a SkipList
func NewNode(key uint64, value [10]byte, height int) *Node {
	return &Node{
		NextNode: make([]*Node, height),
		Value:    value,
		Key:      key,
	}
}

//inserts a new key in the SkipList
func (skipList *SkipList) Insert(node Node) {
	updatePointer := make([]*Node, MaxHeight)
	currentNode := skipList.Head

	for i := skipList.Height; i >= 0; i-- {
		for currentNode.NextNode[i] != nil && currentNode.NextNode[i].Key < node.Key {
			currentNode = currentNode.NextNode[i]
		}
		updatePointer[i] = currentNode
	}

	height := rand.Intn(MaxHeight)

	if height > skipList.Height {
		// Extend updatePointer to accommodate the new height
		newUpdatePointer := make([]*Node, height+1)
		copy(newUpdatePointer, updatePointer)
		for i := skipList.Height + 1; i <= height; i++ {
			newUpdatePointer[i] = skipList.Head
		}
		updatePointer = newUpdatePointer
		skipList.Height = height
	}

	newNode := NewNode(node.Key, node.Value, height)

	for i := 0; i < height; i++ {
		newNode.NextNode[i] = updatePointer[i].NextNode[i]
		updatePointer[i].NextNode[i] = newNode
	}

	skipList.Length++
}



// puts in a new key
func (skipList *SkipList) Put(key uint64, value [10]byte) error {
	newNode :=Node{
		Value: value,
		Key: key,
	}
	skipList.Insert(newNode)


	//limits the size of the kvStore to MaxHeight
if skipList.Length >= MaxLength {
	fmt.Println("SkipList length reached 5, creating SSTable...")
	skipList.CreateSSTable()
}
	//return nil if everything went well
	return nil
}

/*
// returns an exisiting key
func (kvStore myKeyValueStore) Get(key uint64) [10]byte {
	value, keyExists := kvStore.data[key]
	//returns a key only if it exists
	if keyExists {
		return value
	} else {
		var keyNotFound [10]byte
		return keyNotFound
	}
}*/

func (skipList *SkipList) getAllKeys() []uint64{
	keys := make([]uint64, 0, skipList.Length)
	currentNode := skipList.Head.NextNode[0]

	for currentNode != nil {
		keys = append(keys, currentNode.Key)
		currentNode = currentNode.NextNode[0]
	}
	return keys
}

func (skipList *SkipList) find(key uint64) *Node {
	currentNode := skipList.Head

	for i := skipList.Height - 1; i >= 0; i-- {
		for currentNode.NextNode[i] != nil && currentNode.NextNode[i].Key < key {
			currentNode = currentNode.NextNode[i]
		}
	}

	if currentNode.NextNode[0] != nil && currentNode.NextNode[0].Key == key {
		return currentNode.NextNode[0]
	}

	return nil
}

// takes the kvStore, creates an SSTable and saves it in a file
func (skipList *SkipList) CreateSSTable() {
keys := skipList.getAllKeys()
sort.Slice(keys, func(i, j int) bool {
	return keys[i] < keys[j]
})

var builder strings.Builder
for _, key := range keys {
	node := skipList.find(key)
	if node != nil {
		fmt.Fprintf(&builder, "Key: %d, Value: %s\n", key, string(node.Value[:]))
	}
}
	// Write SSTable content to file
	file, err := os.Create("SSTable.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(builder.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("SSTable created successfully.")
}

func (s *SSTable) Format(builder *strings.Builder) {
	builder.WriteString("Min Key: ")
	builder.WriteString(fmt.Sprintf("%d\n", s.minKey))
	builder.WriteString("Max Key: ")
	builder.WriteString(fmt.Sprintf("%d\n", s.maxKey))
	builder.WriteString("Shallow Index: ")
	builder.WriteString(fmt.Sprintf("%v\n", s.shallowIndex))
}

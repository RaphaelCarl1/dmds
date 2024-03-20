package skipList

type SkipList interface {
	Create()
	Insert()
	//Delete()
	Print()
}

type skipList struct {
	head        *Node
	maxHeight   int
	probability float64
}

type Node struct {
	nextNode []*Node
	value    [10]byte
}

func NewNode(nextNode []*Node, height int, value [10]byte) *Node {
	return &Node{
		nextNode: make([]*Node, height),
		value:    value,
	}
}

func NewSkipList(maxHeight int, probability float64) {
	//create new list
	//first node & then pointer to next?
}

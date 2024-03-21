package skipList

type Data interface {
	Create(maxHeigt int, probability float64)
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
	value    int
}

func NewNode(height int, value int) *Node {
	return &Node{
		nextNode: make([]*Node, height),
		value:    value,
	}
}

func Create(maxHeight int, probability float64) *skipList{
	return &skipList{
		head: NewNode(maxHeight, 0),
		maxHeight: maxHeight,
		probability: probability,

	}
}

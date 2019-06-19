package list

type Node struct {
	Value int
	Next  *Node
}

type NodeObject struct {
	Value interface{}
	Next  *NodeObject
}

func NewNodeObject(value interface{}, next *NodeObject) *NodeObject {
	return &NodeObject{
		Value: value,
		Next:  next,
	}
}

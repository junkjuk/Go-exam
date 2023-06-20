package tree

type Tree struct {
	Root *INode
}

type Node struct {
	Data     int
	Children []*INode
}

func (n Node) GetData() int {
	return n.Data
}

func (n Node) GetChildren() []*INode {
	return n.Children
}

func (tree *Tree) Traverse(function func(INode)) {
	traverse(*tree.Root, function)
}

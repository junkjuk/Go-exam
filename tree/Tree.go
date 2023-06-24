package tree

type Tree struct {
	Root Node
}

type Node struct {
	Data     int
	Children []*Node
}

func (n Node) GetData() int {
	return n.Data
}

func (n Node) GetChildren() []INode {
	children := make([]INode, len(n.Children))
	for i, node := range n.Children {
		children[i] = node
	}
	return children
}

func (tree *Tree) Traverse(function func(INode)) {
	traverse(tree.Root, function)
}

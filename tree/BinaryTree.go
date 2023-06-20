package tree

type BinaryTree struct {
	Root *INode
}

type BinaryNode struct {
	Data  int
	left  *INode
	right *INode
}

func (n BinaryNode) GetData() int {
	return n.Data
}

func (n BinaryNode) GetChildren() []*INode {
	return []*INode{n.left, n.right}
}

func (tree *BinaryTree) Traverse(function func(INode)) {
	traverse(*tree.Root, function)
}

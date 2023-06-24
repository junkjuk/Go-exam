package tree

type BinaryTree struct {
	Root *BinaryNode
}

type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

func (n BinaryNode) GetData() int {
	return n.Data
}

func (n BinaryNode) GetChildren() []INode {
	var childrens []INode
	if n.Left != nil {
		childrens = append(childrens, n.Left)
	}
	if n.Right != nil {
		childrens = append(childrens, n.Right)
	}
	return childrens
}

func (tree *BinaryTree) Traverse(function func(INode)) {
	traverse(tree.Root, function)
}

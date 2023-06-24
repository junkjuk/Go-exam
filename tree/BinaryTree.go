package tree

type BinaryTree struct {
	Root *BinaryNode
}

type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

func (n *BinaryNode) SetData(data int) {
	n.Data = data
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

func (tree *BinaryTree) GetBasicNode() INode {
	return &BinaryNode{}
}

func (tree *BinaryTree) Traverse(function func(INode)) {
	traverse(tree.Root, function)
}

func (n *BinaryNode) maxChild() bool {
	if n.Left != nil && n.Right != nil {
		n.Data = max(n.Right.Data, n.Left.Data)
		if !n.Right.maxChild() {
			n.Right = nil
		}
		if !n.Left.maxChild() {
			n.Left = nil
		}
		return true
	}
	if n.Right != nil {
		n.Data = n.Right.Data
		if !n.Right.maxChild() {
			n.Right = nil
		}
		return true
	}
	if n.Left != nil {
		n.Data = n.Left.Data
		if !n.Left.maxChild() {
			n.Left = nil
		}
		return true
	}
	n.Left = nil
	n.Right = nil
	return false
}

func MaxChild(tree BinaryTree) ITree {
	tree.Root.maxChild()
	return &tree
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

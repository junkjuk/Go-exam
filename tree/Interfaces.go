package tree

type ITree interface {
	Traverse(function func(INode))
	GetBasicNode() INode
}

type INode interface {
	GetData() int
	GetChildren() []INode
	SetData(int)
}

func traverse(rootNode INode, function func(INode)) {
	function(rootNode)
	for _, node := range rootNode.GetChildren() {
		traverse(node, function)
	}
}

func SumTree(tree ITree) INode {
	sum := 0
	tree.Traverse(func(node INode) {
		sum += node.GetData()
	})
	node := tree.GetBasicNode()
	node.SetData(sum)
	return node
}

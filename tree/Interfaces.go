package tree

type ITree interface {
	Traverse(function func(INode))
}

type INode interface {
	GetData() int
	GetChildren() []*INode
}

func traverse(rootNode INode, function func(INode)) {
	function(rootNode)
	for _, node := range rootNode.GetChildren() {
		traverse(*node, function)
	}
}

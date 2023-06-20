package main

import (
	"encoding/json"
	"exam/tree"
	"fmt"
)

func main() {
	var n1, n2, n3, n4 tree.INode

	list := tree.Tree{}
	n3 = tree.Node{Data: 3}
	n1 = tree.Node{Data: 1}
	n2 = tree.Node{Data: 2,
		Children: []*tree.INode{&n3},
	}
	n4 = tree.Node{
		Data:     4,
		Children: []*tree.INode{&n2, &n1},
	}
	list.Root = &n4

	js, _ := json.MarshalIndent(list, "", " ")
	jsonStr := string(js)
	fmt.Println(jsonStr)
	sum := 0
	list.Traverse(func(n tree.INode) {
		sum += n.GetData()
	})
	fmt.Println(sum)
}

package main

import (
	"encoding/json"
	"exam/tree"
	"fmt"
)

func main() {
	var n1, n2, n3, n4, n5 tree.BinaryNode

	list := tree.BinaryTree{}
	n2 = tree.BinaryNode{Data: 2}
	n4 = tree.BinaryNode{Data: 4}
	n5 = tree.BinaryNode{Data: 5}
	n3 = tree.BinaryNode{Data: 3, Left: &n4, Right: &n5}
	n1 = tree.BinaryNode{Data: 1, Left: &n2, Right: &n3}

	list.Root = &n1

	js, _ := json.Marshal(list)
	jsonStr := string(js)
	fmt.Println(jsonStr)
	sum := 0
	list.Traverse(func(n tree.INode) {
		sum += n.GetData()
	})
	fmt.Println(sum)
}

package parser

import (
	"encoding/json"
	"exam/tree"
	"fmt"
	"io"
	"strings"
)

func GetTree(in io.Reader) (tree.ITree, error) {
	b, _ := io.ReadAll(in)
	fmt.Println(string(b))
	trees, err := deserializeTree(b)
	if err != nil {
		return nil, err
	}
	return trees, nil
}

func deserializeTree(buffer []byte) (tree.ITree, error) {
	if strings.Contains(string(buffer), "Children") {
		var nTree tree.Tree
		err := json.Unmarshal(buffer, &nTree)
		if err != nil {
			return nil, err
		}
		return &nTree, nil
	} else {
		var nTree tree.BinaryTree
		err := json.Unmarshal(buffer, &nTree)
		if err != nil {
			return nil, err
		}
		return &nTree, nil
	}
}

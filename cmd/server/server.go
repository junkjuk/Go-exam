package main

import (
	"encoding/json"
	"exam/tree"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	h := new(http.ServeMux)

	h.HandleFunc("/sum", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("content-type", "application/json")

		tre, err := getTree(req)
		if err != nil {
			rw.Write([]byte(err.Error()))
			return
		}

		sum := tree.SumTree(tre)
		js, _ := json.Marshal(sum)
		rw.Write(js)
	})

	http.ListenAndServe(":8090", h)
}

func getTree(req *http.Request) (tree.ITree, error) {
	b, _ := io.ReadAll(req.Body)
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

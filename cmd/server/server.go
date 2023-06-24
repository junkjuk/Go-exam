package main

import (
	"encoding/json"
	"exam/cmd/parser"
	"exam/tree"
	"net/http"
)

func main() {
	h := new(http.ServeMux)

	h.HandleFunc("/childmax", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("content-type", "application/json")

		tre, err := parser.GetBinaryTree(req.Body)
		if err != nil {
			_, _ = rw.Write([]byte(err.Error()))
			return
		}

		result := tree.MaxChild(*tre)
		js, _ := json.Marshal(result)
		_, _ = rw.Write(js)
	})

	_ = http.ListenAndServe(":8090", h)
}

package main

import (
	"encoding/json"
	"exam/cmd/parser"
	"exam/tree"
	"net/http"
)

func main() {
	h := new(http.ServeMux)

	h.HandleFunc("/sum", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("content-type", "application/json")

		tre, err := parser.GetTree(req.Body)
		if err != nil {
			_, _ = rw.Write([]byte(err.Error()))
			return
		}

		sum := tree.SumTree(tre)
		js, _ := json.Marshal(sum)
		_, _ = rw.Write(js)
	})

	_ = http.ListenAndServe(":8090", h)
}

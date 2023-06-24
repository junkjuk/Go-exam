package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var url = "http://localhost:8090/childmax"

func main() {
	reader := bufio.NewReader(os.Stdin)
	client := new(http.Client)
	client.Timeout = 10 * time.Second
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Printf("error %s", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error %s", err)
		}
		resp.Body.Close()

		fmt.Println(string(body))
	}
}

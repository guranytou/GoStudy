package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var tasks = []Task{
	{
		ID:      1,
		Title:   "test1",
		Content: "This is test1",
	},
	{
		ID:      2,
		Title:   "test2",
		Content: "This is test2",
	},
}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	http.HandleFunc("/tasks", handler1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

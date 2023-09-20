package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleDefault)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server running on port 4000")
	}
}
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Welcome to http</h1>")
}

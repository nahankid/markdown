package main

import (
	"fmt"
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GenerateMarkdown generates the markdown
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

package main

import (
    "fmt"
    "net/http"
	"github.com/gorilla/mux"
)

func main() {
	rount := mux.NewRouter()

	// ex:http://localhost:80/books/cprogramming/page/1
	// title = cprogramming
	//page = 1
	rount.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })
	
	
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    // })

	// fs := http.FileServer(http.Dir("static/"))
    // http.Handle("/static/", http.StripPrefix
	// ("/static/", fs))

	fmt.Println("Server in runing on port 80")
	http.ListenAndServe(":80", rount)
    // http.ListenAndServe(":80", null)
}
package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "<h1>Book: %s</h1><br><h2>Page: %s</h2>", vars["title"], vars["page"])
	})

	r.HandleFunc("/cal/{num1}/plus/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		num1, _ := strconv.Atoi(vars["num1"])
		num2, _ := strconv.Atoi(vars["num2"])
		fmt.Fprintf(w, "<h1>Calculator: %d Plus = %d</h1>", num1, num2)
		fmt.Fprintf(w, "<h1>Total = %d</h1>", num1+num2)
	})

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "My Name is ping")
    })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	

    http.ListenAndServe(":8080", r)
}
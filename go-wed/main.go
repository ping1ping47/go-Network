package main

import (
    "fmt"
    "net/http"
	"html/template"
)

type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}



func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
	http.ListenAndServe(":80", nil)



	// rount := mux.NewRouter()

	// // ex:http://localhost:80/books/cprogramming/page/1
	// // title = cprogramming
	// //page = 1
	// rount.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    //     vars := mux.Vars(r)
    //     title := vars["title"]
    //     page := vars["page"]

    //     fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    // })
	
	
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    // })

	// fs := http.FileServer(http.Dir("static/"))
    // http.Handle("/static/", http.StripPrefix
	// ("/static/", fs))

	fmt.Println("Server in runing on port 80")
	// http.ListenAndServe(":80", rount)
    // http.ListenAndServe(":80", null)
}
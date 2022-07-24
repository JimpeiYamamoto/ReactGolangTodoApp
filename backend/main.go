package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	title   string
	content string
}

type Info struct {
	todoLst []Todo
	compLst []Todo
}

var info = Info{}

func addTodo(w http.ResponseWriter, r *http.Request) {
	titleqs := r.URL.Query().Get("title")
	contentqs := r.URL.Query().Get("content")
	if titleqs == "" {
		fmt.Fprintf(w, "Title string is missing!")
		return
	}
	if contentqs == "" {
		fmt.Fprintf(w, "Content string is missing!")
		return
	}
	todo := Todo{title: titleqs, content: contentqs}
	info.todoLst = append(info.todoLst, todo)
}

func doneTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/addtodo", addTodo)
	http.HandleFunc("/donetodo", doneTodo)
	http.ListenAndServe(":8080", nil)
}

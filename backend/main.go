package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	indexqs := r.URL.Query().Get("index")
	if indexqs == "" {
		fmt.Fprintf(w, "index is missing!")
		return
	}
	index, err := strconv.ParseInt(indexqs, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid index number!")
		return
	}
	if int(index-1) > len(info.todoLst) {
		fmt.Fprintf(w, "Index out of range info.todoLst!")
		return
	}
	moveTodo := info.todoLst[index]
	info.compLst = append(info.compLst, moveTodo)
	info.todoLst = append(info.todoLst[index:], info.todoLst[index+1:]...)
}

func backTodo(w http.ResponseWriter, r *http.Request) {
	indexqs := r.URL.Query().Get("index")
	if indexqs == "" {
		fmt.Fprintf(w, "index is missing!")
		return
	}
	index, err := strconv.ParseInt(indexqs, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid index number!")
		return
	}
	if int(index-1) > len(info.compLst) {
		fmt.Fprintf(w, "Index out of range info.compLst!")
		return
	}
	moveTodo := info.compLst[int(index)]
	info.compLst = append(info.compLst[:int(index)], info.compLst[int(index)+1:]...)
	info.todoLst = append(info.todoLst, moveTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	indexqs := r.URL.Query().Get("index")
	if indexqs == "" {
		fmt.Fprintf(w, "index is missing!")
		return
	}
	index, err := strconv.ParseInt(indexqs, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid index number!")
		return
	}
	if int(index-1) > len(info.todoLst) {
		fmt.Fprintf(w, "Index out of range info.todoLst!")
		return
	}
	info.todoLst = append(info.todoLst[index:], info.todoLst[index+1:]...)
}

func deleteComp(w http.ResponseWriter, r *http.Request) {
	indexqs := r.URL.Query().Get("index")
	if indexqs == "" {
		fmt.Fprintf(w, "index is missing!")
		return
	}
	index, err := strconv.ParseInt(indexqs, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid index number!")
		return
	}
	if int(index-1) > len(info.compLst) {
		fmt.Fprintf(w, "Index out of range info.compLst!")
		return
	}
	info.compLst = append(info.compLst[:int(index)], info.compLst[int(index)+1:]...)
}

func main() {
	http.HandleFunc("/addtodo", addTodo)
	http.HandleFunc("/donetodo", doneTodo)
	http.HandleFunc("/backtodo", backTodo)
	http.HandleFunc("/deletetodo", deleteTodo)
	http.HandleFunc("/deletecomp", deleteComp)
	http.ListenAndServe(":8080", nil)
}

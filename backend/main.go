package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Todo struct {
	Title   string
	Content string
}

type Info struct {
	TodoLst []Todo
	CompLst []Todo
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
	todo := Todo{Title: titleqs, Content: contentqs}
	info.TodoLst = append(info.TodoLst, todo)
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
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
	if int(index-1) > len(info.TodoLst) {
		fmt.Fprintf(w, "Index out of range info.todoLst!")
		return
	}
	moveTodo := info.TodoLst[index]
	info.CompLst = append(info.CompLst, moveTodo)
	info.TodoLst = append(info.TodoLst[index:], info.TodoLst[index+1:]...)
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
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
	if int(index-1) > len(info.CompLst) {
		fmt.Fprintf(w, "Index out of range info.compLst!")
		return
	}
	moveTodo := info.CompLst[int(index)]
	info.CompLst = append(info.CompLst[:int(index)], info.CompLst[int(index)+1:]...)
	info.TodoLst = append(info.TodoLst, moveTodo)
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
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
	if int(index-1) > len(info.TodoLst) {
		fmt.Fprintf(w, "Index out of range info.todoLst!")
		return
	}
	info.TodoLst = append(info.TodoLst[index:], info.TodoLst[index+1:]...)
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
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
	if int(index-1) > len(info.CompLst) {
		fmt.Fprintf(w, "Index out of range info.compLst!")
		return
	}
	info.CompLst = append(info.CompLst[:int(index)], info.CompLst[int(index)+1:]...)
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "==todo_lst==\n")
	for i := 0; i < len(info.TodoLst); i++ {
		fmt.Fprintf(w, "title: %v, content: %v\n",
			info.TodoLst[i].Title, info.TodoLst[i].Content)
	}
	fmt.Fprintf(w, "==comp_lst==\n")
	for i := 0; i < len(info.CompLst); i++ {
		fmt.Fprintf(w, "title: %v, content: %v\n",
			info.CompLst[i].Title, info.CompLst[i].Content)
	}
	res, err := json.Marshal(info)
	if err != nil {
		fmt.Fprintf(w, "Fail to marshal json!")
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/show", showInfo)
	http.HandleFunc("/addtodo", addTodo)
	http.HandleFunc("/donetodo", doneTodo)
	http.HandleFunc("/backtodo", backTodo)
	http.HandleFunc("/deletetodo", deleteTodo)
	http.HandleFunc("/deletecomp", deleteComp)
	http.ListenAndServe(":8080", nil)
}

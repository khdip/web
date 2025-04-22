package main

import (
	"log"
	"net/http"
	"practice/web/handler"
)

func main() {
	h := handler.GetHandler()
	http.HandleFunc("/", h.GetHome)
	http.HandleFunc("/about", h.GetAbout)
	http.HandleFunc("/json", h.GetJson)
	http.HandleFunc("/todo", h.GetTodo)
	http.HandleFunc("/todo/create", h.CreateTodo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Not Found", err)
	}
}

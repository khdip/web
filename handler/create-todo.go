package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Unable to create ToDo: ", err)
	}
	taskName := r.FormValue("task")
	fmt.Fprintf(w, "Task Name: %s", taskName)
}

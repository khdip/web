package handler

import (
	"encoding/json"
	"net/http"
)

type ToDo struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ToDos := []ToDo{
		{
			Task:        "Learn Go",
			IsCompleted: true,
		},
		{
			Task:        "Get a Project",
			IsCompleted: false,
		},
		{
			Task:        "Become a Pro",
			IsCompleted: false,
		},
	}

	json.NewEncoder(w).Encode(ToDos)
}

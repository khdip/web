package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

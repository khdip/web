package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Want to know about me?")
}

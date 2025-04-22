package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode("{a:1, b:2}")
}

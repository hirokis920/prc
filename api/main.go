package main

import (
	"net/http"
	"strings"
)

// TODO 構造体を作る
func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 4 || parts[2] != "plan" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	planID := parts[3]

	w.Header().Set("Content-Type", "application/json")
	// TODO　レスポンスはダミー
	response := `{"plan_id": "` + planID + `", "status": "active"}`
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/api/plan/{id}", handler)
	http.ListenAndServe(":8080", nil)
}

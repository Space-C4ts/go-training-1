package reverse

import (
	"encoding/json"
	"net/http"
)

type ReverseRequest struct {
	Text string `json:"text"`
}

type ReverseResponse struct {
	Reversed string `json:"reversed"`
}

func Init() {
	http.HandleFunc("/reverse", reverseHandler)
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReverseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Text == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reversed := reverseString(req.Text)

	resp := ReverseResponse{Reversed: reversed}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

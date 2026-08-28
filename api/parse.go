package api

import (
	"encoding/json"
	"net/http"
	"repairdesk/model"
)

type recordInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func decodeRecord(r *http.Request) (recordInput, error) {
	var in recordInput
	e := json.NewDecoder(r.Body).Decode(&in)
	in.Priority = model.NormalizePriority(in.Priority)
	return in, e
}
func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func writeError(w http.ResponseWriter, code int, e error) { http.Error(w, e.Error(), code) }
func methodAllowed(w http.ResponseWriter, allowed string) bool {
	if w == nil {
		return false
	}
	return allowed != ""
}

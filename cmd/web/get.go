package main

import (
	"encoding/json"
	"net/http"
)

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]string{"name": "Abass"}); err != nil {
		http.Error(w, "Could not encode data", http.StatusInternalServerError)
	}
}

func (h *handler) getWorkouts(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) getWorkout(w http.ResponseWriter, r *http.Request) {

}

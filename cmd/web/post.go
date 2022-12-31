package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abassGarane/muscles/domain"
)

func (h *handler) createWorkout(w http.ResponseWriter, r *http.Request) {
	var workout domain.Workout

	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		http.Error(w, fmt.Errorf("Could not decode the request body %v", err).Error(), http.StatusInternalServerError)
		return
	}

	work, err := h.service.CreateWorkout(&workout)
	if err = json.NewEncoder(w).Encode(work); err != nil {
		http.Error(w, "Could not encode the request body", http.StatusInternalServerError)
		return
	}
}

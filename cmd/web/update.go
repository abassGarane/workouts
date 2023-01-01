package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/go-chi/chi/v5"
)

func (h *handler) updateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout domain.Workout

	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		http.Error(w, fmt.Errorf("could not decode the request body %v", err).Error(), http.StatusInternalServerError)
		return
	}
	id := chi.URLParam(r, "id")
	defer r.Body.Close()
	returnedWorkout, err := h.service.UpdateWorkout(id, &workout)
	if err != nil {
		http.Error(w, fmt.Errorf("could not decode the request body %v", err).Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(returnedWorkout); err != nil {
		http.Error(w, fmt.Errorf("could not decode the request body %v", err).Error(), http.StatusInternalServerError)
		return
	}
}

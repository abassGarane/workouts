package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json") // for tests
	if err := json.NewEncoder(w).Encode(map[string]string{"name": "Abass"}); err != nil {
		http.Error(w, "Could not encode data", http.StatusInternalServerError)
	}
}

func (h *handler) getWorkouts(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) getWorkout(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	wkout, err := h.service.GetWorkout(id)
	if err != nil {
		http.Error(w, errors.Wrap(err, "Unable to retrieve workout").Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(wkout)
	if err != nil {
		http.Error(w, errors.Wrap(err, "Unable to encode workout").Error(), http.StatusInternalServerError)
	}
}

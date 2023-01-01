package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/pkg/errors"
)

func (h *handler) createWorkout(w http.ResponseWriter, r *http.Request) {
	var workout domain.Workout

	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		http.Error(w, fmt.Errorf("could not decode the request body %v", err).Error(), http.StatusInternalServerError)
		return
	}

	work, err := h.service.CreateWorkout(&workout)
	if err != nil{
		http.Error(w, errors.Wrap(err,"Could not encode the request body").Error(), http.StatusInternalServerError)
		return	
	}
	if err = json.NewEncoder(w).Encode(work); err != nil {
		http.Error(w, errors.Wrap(err,"Could not encode the request body").Error(), http.StatusInternalServerError)
		return
	}
}

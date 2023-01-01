package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func (h *handler) deleteWorkout(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  defer r.Body.Close()
  err := h.service.DeleteWorkout(id)
  if err != nil{
    http.Error(w, errors.Wrap(err, "Unable to delete workout").Error(), http.StatusInternalServerError)
    return
  }
}

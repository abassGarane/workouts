package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abassGarane/muscles/domain"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	hd.health(w, req)
	res := w.Result()
	require.Equal(t, res.StatusCode, http.StatusOK)
	m := map[string]string{}
	err := json.NewDecoder(res.Body).Decode(&m)
	require.NoError(t, err)
	require.Equal(t, m["name"], "Abass")
	log.Print(res.Header)
	require.Equal(t, res.Header.Get("Content-Type"), "application/json")
}

func TestGetWorkout(t *testing.T) {
	// create a workout
	workout := &domain.Workout{
		ID:   primitive.NewObjectID(),
		Type: "sprints",
		Reps: 40,
		Load: 30,
	}
	body, _ := json.Marshal(workout)
	// require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	hd.createWorkout(w, req)
	result := w.Result()
	defer result.Body.Close()
	wkout := domain.Workout{}
	_ = json.NewDecoder(result.Body).Decode(&wkout)
	log.Print(wkout)

	//get that workout
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", wkout.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	hd.getWorkout(w, req)
	res := w.Result()
	require.Equal(t, res.StatusCode, http.StatusOK)
	retrievedWorkout := domain.Workout{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&retrievedWorkout)
	log.Print(retrievedWorkout)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedWorkout)
	require.Equal(t, retrievedWorkout.Type, "sprints")
	require.Equal(t, retrievedWorkout.Reps, 40)
	require.Equal(t, retrievedWorkout.Load, 30)
	require.NotEmpty(t, retrievedWorkout.ID)

}

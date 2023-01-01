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
		Type: "sprints",
		Reps: 40,
		Load: 30,
	}
	body, err := json.Marshal(workout)
	require.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	hd.createWorkout(w, req)
	result := w.Result()
	defer result.Body.Close()
	err = json.NewDecoder(result.Body).Decode(&workout)
	log.Println("Returned workout", workout)
	require.NoError(t, err)

	//get that workout
	url := fmt.Sprintf("/%s", workout.ID.Hex())
	fmt.Println("url is ::", url)
	req, _ = http.NewRequest(http.MethodGet, "/" + workout.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	w2.Header().Set("Content-Type", "application/json")
	hd.getWorkout(w2, req)
	res := w2.Result()
	fmt.Println(res.Status)
	require.Equal(t, res.StatusCode, http.StatusOK)
	retrievedWorkout := domain.Workout{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&retrievedWorkout)
	log.Printf("Retreived workout:: %#v", retrievedWorkout)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedWorkout)
	require.Equal(t, retrievedWorkout.Type, "sprints")
	require.Equal(t, retrievedWorkout.Reps, 40)
	require.Equal(t, retrievedWorkout.Load, 30)
	require.NotEmpty(t, retrievedWorkout.ID)
}

func TestGetWorkouts(t *testing.T) {

}

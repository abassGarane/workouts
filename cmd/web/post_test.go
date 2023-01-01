package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abassGarane/muscles/domain"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	workout := &domain.Workout{
		Type: "sprints",
		Reps: 40,
		Load: 30,
	}
	body, err := json.Marshal(workout)
	require.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	hd.createWorkout(w, req)
	result := w.Result()
	defer result.Body.Close()
	wkout := domain.Workout{}
	err = json.NewDecoder(result.Body).Decode(&wkout)
	require.NoError(t, err)
	require.NotEmpty(t, wkout)
	require.Equal(t, wkout.Type, "sprints")
	require.Equal(t, wkout.Reps, 40)
	require.Equal(t, wkout.Load, 30)
	require.NotEmpty(t, wkout.ID)
}

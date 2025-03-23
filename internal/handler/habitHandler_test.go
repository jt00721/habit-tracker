package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/jt00721/habit-tracker/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateHabitApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{
			name:     "Valid Create Habit",
			body:     `{"name": "Read a book", "frequency": "daily"}`,
			wantCode: http.StatusCreated,
		},
		{
			name:     "Invalid JSON",
			body:     `{"name": "Missing quote, "frequency": "daily"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid Habit Name",
			body:     `{"name": "", "frequency": "daily"}`,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "Invalid Habit Frequency",
			body:     `{"name": "Valid Habit Name", "frequency": "Not valid frequency"}`,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// payloadBytes, _ := json.Marshal(tt.habitPayload)
			resp, err := http.Post(ts.URL+"/api/habits", "application/json", bytes.NewBufferString(tt.body))
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, resp.StatusCode)

			// body, err := io.ReadAll(resp.Body)
			// assert.NoError(t, err)
			// assert.Equal(t, tt.wantBody, body, string(body))
		})
	}
	// Define habit request payload
	// habitPayload := map[string]string{
	// 	"name":      "Read a book",
	// 	"frequency": "daily",
	// }
	// payloadBytes, _ := json.Marshal(habitPayload)

	// // Make request to create habit
	// resp, err := http.Post(ts.URL+"/api/habits", "application/json", bytes.NewBuffer(payloadBytes))
	// assert.NoError(t, err)
	// assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestGetAllApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	resp, err := http.Get(ts.URL + "/api/habits")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetHabitByIDApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	tests := []struct {
		name     string
		id       string
		wantCode int
	}{
		{
			name:     "Valid Get Habit By ID",
			id:       "1",
			wantCode: http.StatusOK,
		},
		{
			name:     "No Habit ID",
			id:       " ",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid Habit ID",
			id:       "3",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// payloadBytes, _ := json.Marshal(tt.habitPayload)
			resp, err := http.Get(ts.URL + "/api/habits/" + tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, resp.StatusCode)
		})
	}

	// resp, err := http.Get(ts.URL + "/api/habits/1")
	// assert.NoError(t, err)
	// defer resp.Body.Close()

	// assert.Equal(t, http.StatusOK, resp.StatusCode)

	// var habit map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&habit)

	// assert.Equal(t, "Test Habit", habit["Name"])
	// assert.Equal(t, "daily", habit["Frequency"])
}

func TestUpdateHabitApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	createPayload := map[string]string{"name": "Initial", "frequency": "daily"}
	createBytes, _ := json.Marshal(createPayload)
	resp, err := http.Post(ts.URL+"/api/habits", "application/json", bytes.NewBuffer(createBytes))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var created map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&created)
	habitID := int(created["ID"].(float64))
	resp.Body.Close()

	tests := []struct {
		name       string
		id         string
		body       string
		wantStatus int
	}{
		{
			name:       "Valid update",
			id:         fmt.Sprintf("%d", habitID),
			body:       `{"name": "Updated", "frequency": "weekly"}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid habit ID (non-numeric)",
			id:         "abc",
			body:       `{"name": "Oops", "frequency": "daily"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Malformed JSON",
			id:         fmt.Sprintf("%d", habitID),
			body:       `{"name": "Missing quote, "frequency": "daily"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Non-existent habit",
			id:         "99999",
			body:       `{"name": "Ghost", "frequency": "daily"}`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "Invalid Update Habit Name",
			id:         fmt.Sprintf("%d", habitID),
			body:       `{"name": "", "frequency": "weekly"}`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "Invalid Update Habit Frequency",
			id:         fmt.Sprintf("%d", habitID),
			body:       `{"name": "Updated Habit", "frequency": "Not valid frequency"}`,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/api/habits/%s", ts.URL, tt.id)
			req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}

	// resp, err := http.NewRequest(http.MethodPut, ts.URL+"/api/habits/1", bytes.NewBuffer(updateBytes))
	// assert.NoError(t, err)
	// resp.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// updateResp, err := client.Do(resp)
	// assert.NoError(t, err)
	// defer updateResp.Body.Close()

	// assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	// var updatedHabit map[string]interface{}
	// json.NewDecoder(updateResp.Body).Decode(&updatedHabit)

	// assert.Equal(t, "Drink more water", updatedHabit["Name"])
	// assert.Equal(t, "weekly", updatedHabit["Frequency"])
}

func TestDeleteHabitApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	tests := []struct {
		name     string
		id       string
		wantCode int
	}{
		{
			name:     "Valid Delete Habit",
			id:       "1",
			wantCode: http.StatusOK,
		},
		{
			name:     "No Habit ID",
			id:       " ",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid Habit ID",
			id:       "3",
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// payloadBytes, _ := json.Marshal(tt.habitPayload)
			req, err := http.NewRequest(http.MethodDelete, ts.URL+"/api/habits/"+tt.id, nil)
			assert.NoError(t, err)

			client := &http.Client{}
			resp, err := client.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantCode, resp.StatusCode)
		})
	}

	// resp, err := http.NewRequest(http.MethodDelete, ts.URL+"/api/habits/1", nil)
	// assert.NoError(t, err)

	// client := &http.Client{}
	// deleteResp, err := client.Do(resp)
	// assert.NoError(t, err)
	// defer deleteResp.Body.Close()

	// assert.Equal(t, http.StatusOK, deleteResp.StatusCode)

	// body, err := io.ReadAll(deleteResp.Body)
	// assert.NoError(t, err)

	// assert.Contains(t, string(body), "Habit deleted")
}

func TestMarkHabitCompletedApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	tests := []struct {
		name     string
		id       string
		wantCode int
	}{
		{
			name:     "Valid Mark Habit Completed",
			id:       "1",
			wantCode: http.StatusOK,
		},
		{
			name:     "No Habit ID",
			id:       " ",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid Habit ID",
			id:       "3",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// payloadBytes, _ := json.Marshal(tt.habitPayload)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+"/api/habits/"+tt.id+"/mark_complete", nil)
			assert.NoError(t, err)

			client := &http.Client{}
			resp, err := client.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantCode, resp.StatusCode)
		})
	}

	// resp, err := http.NewRequest(http.MethodPatch, ts.URL+"/api/habits/1/mark_complete", nil)
	// assert.NoError(t, err)
	// // defer resp.Body.Close()

	// client := &http.Client{}
	// patchresp, err := client.Do(resp)
	// assert.NoError(t, err)
	// defer patchresp.Body.Close()

	// assert.Equal(t, http.StatusOK, patchresp.StatusCode)

	// body, err := io.ReadAll(patchresp.Body)
	// assert.NoError(t, err)

	// assert.Contains(t, string(body), "Habit Completed", body)

	// getHabitByIdResp, err := http.Get(ts.URL + "/api/habits/1")
	// assert.NoError(t, err)
	// defer getHabitByIdResp.Body.Close()

	// assert.Equal(t, http.StatusOK, getHabitByIdResp.StatusCode)

	// var habit map[string]interface{}
	// json.NewDecoder(getHabitByIdResp.Body).Decode(&habit)

	// assert.Equal(t, float64(6), habit["CurrentStreak"], habit)
	// assert.Equal(t, float64(11), habit["TotalCompletions"], habit)
}

func TestGetStreaksApi(t *testing.T) {
	ts, _, teardown := testutils.NewTestServer(t)
	defer teardown()

	habitNoStreak := map[string]interface{}{
		"name":           "Sleep early",
		"frequency":      "daily",
		"current_streak": 0,
	}
	payload2, _ := json.Marshal(habitNoStreak)
	http.Post(ts.URL+"/api/habits", "application/json", bytes.NewBuffer(payload2))

	tests := []struct {
		name       string
		wantStatus int
	}{
		{
			name:       "Get streaks with results",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + "/api/habits/streaks")
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}

	// resp, err := http.Get(ts.URL + "/api/habits/streaks")
	// assert.NoError(t, err)
	// assert.Equal(t, http.StatusOK, resp.StatusCode)

	// var habits []map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&habits)

	// assert.Equal(t, "Test Habit", habits[0]["Name"], habits)
	// assert.Equal(t, "daily", habits[0]["Frequency"], habits)
}

package handlers

import (
	"net/http"
	"time"
)

type Health struct {
	Status  string `json:"status"`
	Healthy bool   `json:"healthy"`
}

type FullHealth struct {
	Status    string    `json:"status"`
	Healthy   bool      `json:"healthy"`
	Timestamp time.Time `json:"timestamp"`
	Database  bool      `json:"database"`
}

func HealthHandler(res http.ResponseWriter, req *http.Request) {
	if err := EncodeJsonBody(res, Health{
		Status:  "pass",
		Healthy: true,
	}); err != nil {
		return
	}
}

func FullHealthHandler(res http.ResponseWriter, req *http.Request) {
	st, err := GetStore(req)
	database := err == nil && st.HealthCheck()

	healthy := database

	status := "pass"
	if !healthy {
		status = "failed"
	}

	fullHealth := FullHealth{
		Status:    status,
		Healthy:   healthy,
		Database:  database,
		Timestamp: time.Now(),
	}

	// Encode the created user response
	if err := EncodeJsonBody(res, fullHealth); err != nil {
		return
	}
}

package models

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Cause      string `json:"cause"`
}

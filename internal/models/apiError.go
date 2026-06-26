package models

type APIErrorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIError struct {
	Success bool `json:"success"`
	error   APIErrorDetails
}

package utils

import "todo-app/internal/models"

func CreateError(statusCode int, message string, cause string) *models.AppError {
	return &models.AppError{
		Message:    message,
		StatusCode: statusCode,
		Cause:      cause,
	}
}

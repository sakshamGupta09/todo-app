package utils

import (
	"net/http"
	"strconv"
)

func ToInt(r *http.Request, key string) (int, error) {
	value := r.PathValue(key)
	return strconv.Atoi(value)
}

package e

import "net/http"

const (
	SUCCESS       = http.StatusOK
	InvalidParams = http.StatusBadRequest
	ERROR         = http.StatusInternalServerError
)
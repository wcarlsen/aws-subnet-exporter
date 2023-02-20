package utils

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("healthy"))
}

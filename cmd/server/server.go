package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/szykol/yang-validator/pkg/domain"
	"github.com/szykol/yang-validator/pkg/validator"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	input := domain.ValidationInput{Contents: bytes}

	if err := validator.ValidateYang(input); err != nil {
		http.Error(w, fmt.Errorf("Error: %w", err).Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		slog.Error("error handling request", "error", err)
	}
}

func main() {
	http.HandleFunc("/validate", validateHandler)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		slog.Error("error serving http", "error", err)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func processError(err error, w http.ResponseWriter) {
	var (
		syn    *json.SyntaxError
		typ    *json.UnmarshalTypeError
		max    *http.MaxBytesError
		valErr validator.ValidationErrors
	)

	switch {
	case errors.As(err, &valErr):
		http.Error(w, fmt.Sprintf("validation error: %s", valErr.Error()), http.StatusBadRequest)
	case errors.As(err, &syn):
		http.Error(w, fmt.Sprintf("invalid json at offset %d", syn.Offset), http.StatusBadRequest)
	case errors.As(err, &typ):
		http.Error(w, fmt.Sprintf("field %q: expected %s, got %s", typ.Field, typ.Type, typ.Value), http.StatusBadRequest)
	case errors.Is(err, io.EOF):
		http.Error(w, "empty body", http.StatusBadRequest)
	case errors.As(err, &max):
		http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
	default:
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}
}

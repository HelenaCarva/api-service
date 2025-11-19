package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// JSONResponse represents a standard JSON response structure.
type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// WriteJSON writes a JSON response with the provided status code.
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	return err
}

// ReadJSON reads JSON from the request body into the target struct.
func ReadJSON(w http.ResponseWriter, r *http.Request, target interface{}) error {
	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(target)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return errors.New("malformed JSON at position " + string(rune(syntaxError.Offset)))
		case errors.Is(err, errors.New("http: request body too large")):
			return errors.New("request body must not be larger than 1MB")
		case errors.As(err, &unmarshalTypeError):
			return errors.New("invalid value for field " + unmarshalTypeError.Field)
		case errors.Is(err, errors.New("json: unknown field")):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return errors.New("unknown field " + fieldName)
		case errors.As(err, &invalidUnmarshalError):
			return errors.New("invalid target type for unmarshalling")
		default:
			return err
		}
	}

	return nil
}

// ErrorResponse returns a JSON error response.
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	resp := JSONResponse{
		Success: false,
		Error:   message,
	}
	_ = WriteJSON(w, status, resp)
}

// SuccessResponse returns a JSON success response.
func SuccessResponse(w http.ResponseWriter, status int, data interface{}) {
	resp := JSONResponse{
		Success: true,
		Data:    data,
	}
	_ = WriteJSON(w, status, resp)
}
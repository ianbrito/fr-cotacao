package dto

import "net/http"

type ErrorResponse struct {
	Message string   `json:"message"`
	Status  int      `json:"-"`
	Errors  []string `json:"errors"`
}

func NewErrorResponse(message string, status int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status:  status,
	}
}

func (err *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	if err.Errors == nil {
		err.Errors = []string{}
	}
	return nil
}

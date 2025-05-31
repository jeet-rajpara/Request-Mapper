package error

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	ErrCode        string       `json:"code"`
	HttpStatusCode int          `json:"httpStatusCode,omitempty"`
	ErrMessage     string       `json:"message"`
	PayLoad        *interface{} `json:"invalidParameters,omitempty"`
}

// generateErrorCodeAndMessage generates error code and message
func GenerateErrorCodeAndMessage(httpErrCode int, msg string) *Error {
	return &Error{
		HttpStatusCode: httpErrCode,
		ErrMessage:     msg,
	}
}

func (e *Error) Error() string {
	return e.ErrMessage
}

// errro response
func ErrorResponse(w http.ResponseWriter, errCode int, msg string) {
	err := GenerateErrorCodeAndMessage(errCode, msg)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
}

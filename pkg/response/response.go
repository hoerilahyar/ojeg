package response

import (
	"encoding/json"
	"net/http"

	"ojeg/pkg/errors"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"` // app-specific code (e.g. 0 = OK, 701 = error)
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Success sends a success response
func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BaseResponse{
		Status: "success",
		Code:   0,
		Data:   data,
	})
}

// Warning sends a warning response (e.g. soft validations, no error)
func Warning(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BaseResponse{
		Status: "warning",
		Code:   code,
		Error: map[string]interface{}{
			"message": message,
		},
	})
}

// Error sends an error response using your AppError
func Error(w http.ResponseWriter, err interface{}) {
	var res ErrorResponse

	res.Status = "error"

	switch e := err.(type) {
	case *errors.AppError:
		res.Code = e.Code
		res.Message = e.Message
		w.WriteHeader(e.HTTPStatus)
	case string:
		res.Code = 700
		res.Message = e
		w.WriteHeader(http.StatusBadRequest)
	default:
		res.Code = 999
		res.Message = "Internal server error"
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(res)
}

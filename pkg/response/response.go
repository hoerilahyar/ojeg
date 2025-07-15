package response

import (
	"encoding/json"
	"net/http"

	"ojeg/pkg/errors"
)

type BaseResponse struct {
	Status string      `json:"status"` // success, error, warning
	Code   int         `json:"code"`   // app-specific code (e.g. 0 = OK, 701 = error)
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
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
func Error(w http.ResponseWriter, appErr *errors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.HTTPStatus)
	json.NewEncoder(w).Encode(BaseResponse{
		Status: "error",
		Code:   appErr.Code,
		Error: map[string]interface{}{
			"message": appErr.Message,
		},
	})
}

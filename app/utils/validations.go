package utils

import "github.com/revel/revel"

// ValidationResponse : simple struct to store the Message & Key of a validation error for response
type ValidationResponse struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

// ErrorMapToArray : convert ErrorMap to Array
func ErrorMapToArray(errorMap map[string]*revel.ValidationError) []ValidationResponse {
	errors := []ValidationResponse{}
	for _, v := range errorMap {
		errors = append(errors, ValidationResponse{Key: v.Key, Message: v.Message})
	}
	return errors
}

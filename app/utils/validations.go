package utils

import "github.com/revel/revel"

type ValidationResponse struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func ErrorMapToArray(errorMap map[string]*revel.ValidationError) []ValidationResponse {
	errors := []ValidationResponse{}
	for _, v := range errorMap {
		errors = append(errors, ValidationResponse{Key: v.Key, Message: v.Message})
	}
	return errors
}

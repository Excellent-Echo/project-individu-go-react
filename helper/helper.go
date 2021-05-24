package helper

import (
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// helper function for response API
func APIResponse(msg string, code int, status string, data interface{}) Response {
	var meta = Meta{
		Message: msg,
		Code:    code,
		Status:  status,
	}

	var response = Response{
		Meta: meta,
		Data: data,
	}
	return response
}

// helper function for split error, better looks error information
func SplitErrorInformation(err error) []string {
	var errors []string

	for _, er := range err.(validator.ValidationErrors) {
		errors = append(errors, er.Error())
	}
	return errors
}

// helper function for validate id number
func ValidateIDNumber(id string) error {
	if num, err := strconv.Atoi(id); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input must be a valid id user")
	}
	return nil
}

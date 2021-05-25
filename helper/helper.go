package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"math"
	"strconv"
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

func SplitErrorInformation(err error) []string {
	var errorInfo []string

	for _, e := range err.(validator.ValidationErrors) {
		errorInfo = append(errorInfo, e.Error())
	}

	return errorInfo
}

func ValidateIDNumber(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input must be valid id user")
	}
	return nil
}

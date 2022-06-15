package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitErr := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
		Error:   splitErr,
	}
}
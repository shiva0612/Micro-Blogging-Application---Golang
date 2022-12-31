package errs

import (
	"net/http"
)

type AppErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewTechErr() *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error, please try again later...",
	}
}

func NewUserErr(code int, msg string) *AppErr {
	return &AppErr{
		Code:    code,
		Message: msg,
	}
}

func NewUnAuthorizedErr() *AppErr {
	return &AppErr{
		Code:    http.StatusUnauthorized,
		Message: "UnAuthorized",
	}
}
func NewUserNotFoundErr() *AppErr {
	return &AppErr{
		Code:    http.StatusBadRequest,
		Message: "user not found",	}
}

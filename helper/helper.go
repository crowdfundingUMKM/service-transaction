package helper

import (
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

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())

	}

	return errors
}

// Admin Request
type UserInvestor struct {
	UnixInvestor          string `json:"unix_investor"`
	StatusAccountInvestor string `json:"status_account_investor"`
}

type InvestorStatusResponse struct {
	Meta Meta         `json:"meta"`
	Data UserInvestor `json:"data"`
}

type VerifyTokenApiInvestorResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		UnixInvestor string `json:"investor_id"`
		Succes       string `json:"success"`
	} `json:"data"`
}

package api

import (
	"errors"
	"os"
)

// check service user investor
func CheckServiceUserInvestor() error {
	serviceAdmin := os.Getenv("SERVICE_INVESTOR")
	if serviceAdmin == "" {
		return errors.New("service user investor is empty")
	}
	return nil
}

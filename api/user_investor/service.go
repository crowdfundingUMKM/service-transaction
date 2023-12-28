package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceInvestor() error {
	serviceAdmin := os.Getenv("SERVICE_INVESTOR_HOST")
	if serviceAdmin == "" {
		return errors.New("service investor is empty")
	}
	return nil
}

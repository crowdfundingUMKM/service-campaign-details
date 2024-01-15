package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceUserInvestor() error {
	serviceAdmin := os.Getenv("SERVICE_CAMPAIGN")
	if serviceAdmin == "" {
		return errors.New("service admin is empty")
	}
	return nil
}

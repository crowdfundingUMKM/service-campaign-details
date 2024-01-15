package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceUserCampaign() error {
	serviceAdmin := os.Getenv("SERVICE_ADMIN")
	if serviceAdmin == "" {
		return errors.New("service admin is empty")
	}
	return nil
}

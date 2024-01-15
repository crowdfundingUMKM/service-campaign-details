package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceUserCampaign() error {
	serviceAdmin := os.Getenv("SERVICE_CAMPAIGN")
	if serviceAdmin == "" {
		return errors.New("service user campaign is empty")
	}
	return nil
}

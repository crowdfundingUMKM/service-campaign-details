package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-campaign-details/helper"
)

func GetCampaignId(input CampaignIdInput) (string, error) {
	// check service admin

	err := CheckServiceUserCampaign()
	if err != nil {
		return "", err
	}

	campaignID := helper.UserAdmin{}
	campaignID.UnixAdmin = input.UnixID
	// fetch get /getAdminID from service api
	serviceAdmin := os.Getenv("SERVICE_CAMPAIGN")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return campaignID.UnixAdmin, errors.New("service admin is empty")
	}
	resp, err := http.Get(serviceAdmin + "/api/v1/campaign/getUserCampaignID/" + campaignID.UnixAdmin)
	if err != nil {
		return campaignID.UnixAdmin, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return campaignID.UnixAdmin, errors.New("failed to get admin status or admin not found")
	}

	var response helper.AdminStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	} else if response.Data.StatusAccountAdmin == "deactive" {
		return "", errors.New("admin account is deactive")
	} else if response.Data.StatusAccountAdmin == "active" {
		return campaignID.UnixAdmin, nil
	} else {
		return "", errors.New("invalid admin status")
	}
}

// verify token from service user admin
func VerifyTokenAdmin(input string) (string, error) {

	err := CheckServiceUserCampaign()
	if err != nil {
		return "", err
	}

	// fetch get /verifyToken from service api
	serviceAdmin := os.Getenv("SERVICE_ADMIN")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return "", errors.New("service user campaign is empty")
	}
	req, err := http.NewRequest("GET", serviceAdmin+"/api/v1/verifyTokenCampaign", nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+input)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid token, account deactive or token expired")
	}

	var response helper.VerifyTokenApiUserCampaignResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	}

	return response.Data.UnixCampaign, nil

}
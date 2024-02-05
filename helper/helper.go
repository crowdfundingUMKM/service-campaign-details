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
type UserAdmin struct {
	UnixAdmin          string `json:"unix_admin"`
	StatusAccountAdmin string `json:"status_account_admin"`
}

type AdminStatusResponse struct {
	Meta Meta      `json:"meta"`
	Data UserAdmin `json:"data"`
}

type VerifyTokenApiAdminResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		UnixAdmin string `json:"admin_id"`
		Succes    string `json:"success"`
	} `json:"data"`
}

// User Investor Request
type UserInvestor struct {
	UnixInvestor          string `json:"unix_investor"`
	StatusAccountInvestor string `json:"status_account_investor"`
}

type UserInvestorResponse struct {
	Meta Meta         `json:"meta"`
	Data UserInvestor `json:"data"`
}

type VerifyTokenApiUserInvestorResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		UnixInvestor string `json:"investor_id"`
		Succes       string `json:"success"`
	} `json:"data"`
}

// User Campaign Request
type UserCampaign struct {
	UnixCampaign          string `json:"unix_campaign"`
	StatusAccountCampaign string `json:"status_account_campaign"`
}

type UserCampaignResponse struct {
	Meta Meta         `json:"meta"`
	Data UserCampaign `json:"data"`
}

type VerifyTokenApiUserCampaignResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		UnixCampaign string `json:"user_campaign_id"`
		Succes       string `json:"success"`
	} `json:"data"`
}

// User Reviewer Request
type UserReviwer struct {
	UnixReviewer          string `json:"unix_reviewer"`
	StatusAccountReviewer string `json:"status_account_reviewer"`
}

type UserReviewerResponse struct {
	Meta Meta        `json:"meta"`
	Data UserReviwer `json:"data"`
}

type VerifyTokenApiUserReviewerResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		UnixReviewer string `json:"reviewer_id"`
		Succes       string `json:"success"`
	} `json:"data"`
}

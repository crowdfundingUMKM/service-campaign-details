package core

import (
	"time"
)

// CampaignDetail represents the CampaignDetails table entity
type CampaignDetail struct {
	ID               int       `json:"id"`
	UnixID           string    `json:"unix_id"`
	UserCampaignID   string    `json:"user_campaign_id"`
	NameCampaign     string    `json:"name_campaign"`
	Category         string    `json:"category"`
	Description      string    `json:"description"`
	AddressCampaign  string    `json:"address_campaign"`
	GoalAmount       string    `json:"goal_amount"`
	CurrentAmount    string    `json:"current_amount"`
	MinimumInvest    string    `json:"minimum_invest"`
	InterestRate     int       `json:"interest_rate"`
	TenorPeriod      int       `json:"tenor_period"`
	DeadlineCampaign time.Time `json:"deadline_campaign"`
	BusinessProposal string    `json:"business_proposal"`
	Perks            string    `json:"perks"`
	LinkBisnis       string    `json:"link_bisnis"`
	StatusCampaign   string    `json:"status_campaign"`
	Slug             string    `json:"slug"`
	Rating           int       `json:"rating"`
	DoneCampaign     string    `json:"done_campaign"`
	Refund           string    `json:"refund"`
	StatusRefund     string    `json:"status_refund"`
	UpdateIDReviewer string    `json:"updateId_reviewer"`
	UpdateIDAdmin    string    `json:"updateId_admin"`
	UpdateAtAdmin    time.Time `json:"updateAt_admin"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CampaignImage represents the CampaignImages table entity
type CampaignImage struct {
	ID         int       `json:"id"`
	CampaignID string    `json:"campaign_id"`
	FileName   string    `json:"file_name"`
	IsPrimary  int       `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ApplicationFund represents the ApplicationFunds table entity
type ApplicationFund struct {
	ID               int       `json:"id"`
	CampaignID       string    `json:"campaign_id"`
	StatusSubmission string    `json:"status_submission"`
	StatusAdminID    string    `json:"status_adminId"`
	SubmissionAt     time.Time `json:"submission_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

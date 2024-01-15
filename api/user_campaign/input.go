package api

type CampaignIdInput struct {
	UnixID string `uri:"campaign_id" binding:"required"`
}

type VerifyTokenUserCampaignInput struct {
	Token string `json:"token" binding:"required"`
}

type CampaignId struct {
	UnixCampaign string
}

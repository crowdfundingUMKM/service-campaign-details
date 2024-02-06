package core

type CampaignDetailFormatter struct {
	ID              int    `json:"id"`
	UnixID          string `json:"unix_id"`
	UserCampaignID  string `json:"user_campaign_id"`
	TitleUMKM       string `json:"title_umkm"`
	CategoryUMKM    string `json:"category_umkm"`
	DescriptionUMKM string `json:"description_umkm"`
	AddressUMKM     string `json:"address_umkm"`
}

func FormaterCampaignDetail(campaign CampaignDetail, userID string) CampaignDetailFormatter {
	return CampaignDetailFormatter{
		ID:              campaign.ID,
		UnixID:          campaign.UnixID,
		UserCampaignID:  userID,
		TitleUMKM:       campaign.TitleUMKM,
		CategoryUMKM:    campaign.CategoryUMKM,
		DescriptionUMKM: campaign.DescriptionUMKM,
		AddressUMKM:     campaign.AddressUMKM,
	}
}

type CampaignInforamtioUMKMFormatter struct {
	GoalAmount       int    `json:"goal_amount"`
	MinimumInvest    int    `json:"minimum_invest"`
	InterestRate     int    `json:"interest_rate"`
	TenorPeriod      int    `json:"tenor_period"`
	DeadlineCampaign string `json:"deadline_campaign"`
}

func FormaterInformationUMKM(campaignDetail CampaignDetail) CampaignInforamtioUMKMFormatter {
	return CampaignInforamtioUMKMFormatter{
		GoalAmount:       campaignDetail.GoalAmount,
		MinimumInvest:    campaignDetail.MinimumInvest,
		InterestRate:     campaignDetail.InterestRate,
		TenorPeriod:      campaignDetail.TenorPeriod,
		DeadlineCampaign: campaignDetail.DeadlineCampaign.String(),
	}
}

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

func FormaterCampaignDetail(campaign CampaignDetail, userID string) CampaignDetail {
	return CampaignDetail{
		ID:              campaign.ID,
		UnixID:          campaign.UnixID,
		UserCampaignID:  userID,
		TitleUMKM:       campaign.TitleUMKM,
		CategoryUMKM:    campaign.CategoryUMKM,
		DescriptionUMKM: campaign.DescriptionUMKM,
		AddressUMKM:     campaign.AddressUMKM,
	}
}

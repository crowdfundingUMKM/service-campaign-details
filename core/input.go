package core

import "time"

type CreateCampaignInput struct {
	TitleUMKM       string `json:"title_umkm" form:"title" binding:"required"`
	CategoryUMKM    string `json:"category_umkm" form:"category" binding:"required"`
	DescriptionUMKM string `json:"description_umkm" form:"description" binding:"required"`
	AddressUMKM     string `json:"address_umkm" form:"address" binding:"required"`
}

type UpdateInformationInput struct {
	GoalAmount       int       `json:"goal_amount" form:"goal_amount" `
	MinimumInvest    int       `json:"minimum_invest" form:"minimum_invest" `
	InterestRate     int       `json:"interest_rate" form:"interest_rate" `
	TenorPeriod      int       `json:"tenor_period" form:"tenor_period" `
	DeadlineCampaign time.Time `json:"deadline_campaign" form:"deadline_campaign"`
}

type UpdateMoreInformationInput struct {
	Benefit      string `json:"benefit" form:"benefit" binding:"required"`
	LinkBusiness string `json:"link_business" form:"link_business" binding:"required"`
}

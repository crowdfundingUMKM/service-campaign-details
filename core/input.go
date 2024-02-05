package core

type CreateCampaignInput struct {
	TitleUMKM       string `json:"title_umkm" form:"title" binding:"required"`
	CategoryUMKM    string `json:"category_umkm" form:"category" binding:"required"`
	DescriptionUMKM string `json:"description_umkm" form:"description" binding:"required"`
	AddressUMKM     string `json:"address_umkm" form:"address" binding:"required"`
}

type CreateInformationInput struct {
	GoalAmount       int    `json:"goal_amount" form:"goal_amount" binding:"required"`
	MinimumInvest    int    `json:"minimum_invest" form:"minimum_invest" binding:"required"`
	InterestRate     int    `json:"interest_rate" form:"interest_rate" binding:"required"`
	TenorPeriod      int    `json:"tenor_period" form:"tenor_period" binding:"required"`
	DeadlineCampaign string `json:"deadline_campaign" form:"deadline_campaign" binding:"required"`
}

type CreateMoreInformationInput struct {
	Benefit      string `json:"benefit" form:"benefit" binding:"required"`
	LinkBusiness string `json:"link_business" form:"link_business" binding:"required"`
}

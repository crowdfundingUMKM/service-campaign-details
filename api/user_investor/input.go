package api

type InvestorIdInput struct {
	UnixID string `uri:"investor_id" binding:"required"`
}

type VerifyTokenUserInvestorInput struct {
	Token string `json:"token" binding:"required"`
}

type InvestorId struct {
	UnixInvestor string
}

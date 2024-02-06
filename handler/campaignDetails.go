package handler

import (
	"net/http"
	api_campaign "service-campaign-details/api/user_campaign"
	"service-campaign-details/auth"
	"service-campaign-details/core"
	"service-campaign-details/helper"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

type campaignDetailHandler struct {
	campaignDetailService core.Service
	authService           auth.Service
}

var (
	storageClient *storage.Client
)

func NewCampaignDetailsHandler(campaignDetailService core.Service, authService auth.Service) *campaignDetailHandler {
	return &campaignDetailHandler{campaignDetailService, authService}
}

func (h *campaignDetailHandler) CreateCampaign(c *gin.Context) {
	// input from user
	// mapping input to struct
	// call service
	// response
	var input core.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUserCampaign := c.MustGet("currentUserCampaign").(api_campaign.CampaignId)
	// check if user_campaign_id is now db, return youhave campaign rightnow
	campaignExist, err := h.campaignDetailService.FindCampaignByCampaignId(currentUserCampaign.UnixCampaign)
	if err != nil {
		response := helper.APIResponse("Create campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignExist.UnixID != "" {
		// User already has a campaign, return an error
		response := helper.APIResponse("Create campaign failed - User already has a campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// craete campaign
	createCampaign, err := h.campaignDetailService.CreateCampaign(input, currentUserCampaign.UnixCampaign)

	if err != nil {
		response := helper.APIResponse("Create campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormaterCampaignDetail(createCampaign, currentUserCampaign.UnixCampaign)
	response := helper.APIResponse("Create campaign success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignDetailHandler) UpdateInformationUmkm(c *gin.Context) {
	// input from user
	// mapping input to struct
	// call service
	// response
	var input core.UpdateInformationInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update information umkm failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUserCampaign := c.MustGet("currentUserCampaign").(api_campaign.CampaignId)

	// check if user_campaign_id is now db, return youhave campaign rightnow
	campaignExist, err := h.campaignDetailService.FindCampaignByCampaignId(currentUserCampaign.UnixCampaign)
	if err != nil {
		response := helper.APIResponse("Update information umkm failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignExist.UnixID == "" {
		// User already has a campaign, return an error
		response := helper.APIResponse("Update information umkm failed - User doesn't have a campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// update information umkm
	updateInformationUmkm, err := h.campaignDetailService.UpdateInformationUmkm(input, campaignExist.UnixID)

	if err != nil {
		response := helper.APIResponse("Update information umkm failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := core.FormaterInformationUMKM(updateInformationUmkm)
	response := helper.APIResponse("Update information umkm success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

package core

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	FindCampaignByCampaignId(campaignId string) (CampaignDetail, error)
	CreateCampaign(input CreateCampaignInput, userId string) (CampaignDetail, error)
	UpdateInformationUmkm(input UpdateInformationInput, userId string) (CampaignDetail, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaignByCampaignId(campaignId string) (CampaignDetail, error) {
	campaign, err := s.repository.FindByUserID(campaignId)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput, userId string) (CampaignDetail, error) {
	CampaignDetail := CampaignDetail{}
	CampaignDetail.UnixID = uuid.New().String()[:12]
	CampaignDetail.UserCampaignID = userId
	CampaignDetail.TitleUMKM = input.TitleUMKM
	CampaignDetail.CategoryUMKM = input.CategoryUMKM
	CampaignDetail.DescriptionUMKM = input.DescriptionUMKM
	CampaignDetail.AddressUMKM = input.AddressUMKM
	CampaignDetail.StatusCampaign = "deactive"
	CampaignDetail.CreatedAt = time.Now()

	newCampaign, err := s.repository.SaveCampaign(CampaignDetail)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil

}

func (s *service) UpdateInformationUmkm(input UpdateInformationInput, campaignId string) (CampaignDetail, error) {
	campaign, err := s.repository.FindByUnixID(campaignId)
	if err != nil {
		return campaign, err
	}
	campaign.GoalAmount = input.GoalAmount
	campaign.MinimumInvest = input.MinimumInvest
	campaign.InterestRate = input.InterestRate
	campaign.TenorPeriod = input.TenorPeriod
	campaign.DeadlineCampaign = input.DeadlineCampaign

	updatedCampaign, err := s.repository.UpdateInformationCampaign(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

package core

import "gorm.io/gorm"

type Repository interface {
	SaveCampaign(campaign CampaignDetail) (CampaignDetail, error)
	FindByUnixID(unix_id string) (CampaignDetail, error)
	FindByUserID(user_id string) (CampaignDetail, error)

	UpdateInformationCampaign(campaign CampaignDetail) (CampaignDetail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveCampaign(campaignDetail CampaignDetail) (CampaignDetail, error) {
	err := r.db.Create(&campaignDetail).Error
	if err != nil {
		return campaignDetail, err
	}
	return campaignDetail, nil
}

func (r *repository) FindByUnixID(unix_id string) (CampaignDetail, error) {
	var campaignDetail CampaignDetail

	err := r.db.Where("unix_id = ?", unix_id).Find(&campaignDetail).Error

	if err != nil {
		return campaignDetail, err
	}
	return campaignDetail, nil

}

func (r *repository) FindByUserID(user_id string) (CampaignDetail, error) {
	var campaignDetail CampaignDetail

	err := r.db.Where("user_campaign_id = ?", user_id).Find(&campaignDetail).Error

	if err != nil {
		return campaignDetail, err
	}
	return campaignDetail, nil
}

func (r *repository) UpdateInformationCampaign(campaignDetail CampaignDetail) (CampaignDetail, error) {
	err := r.db.Model(&campaignDetail).Where("user_campaign_id = ?", campaignDetail.UserCampaignID).Updates(CampaignDetail{GoalAmount: campaignDetail.GoalAmount, MinimumInvest: campaignDetail.MinimumInvest, InterestRate: campaignDetail.InterestRate, TenorPeriod: campaignDetail.TenorPeriod, DeadlineCampaign: campaignDetail.DeadlineCampaign}).Error
	if err != nil {
		return campaignDetail, err
	}
	return campaignDetail, nil
}

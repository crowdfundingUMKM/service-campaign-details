package main

import (
	"fmt"
	"log"
	"os"
	"service-campaign-details/auth"
	"service-campaign-details/config"
	"service-campaign-details/core"
	"service-campaign-details/database"
	"service-campaign-details/handler"
	"service-campaign-details/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Service Campaign Details")
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log
	// config.InitLog()

	// setup repository
	db := database.NewConnectionDB()
	campaignDetailRepository := core.NewRepository(db)

	// SETUP SERVICE
	campaignDetailService := core.NewService(campaignDetailRepository)
	authService := auth.NewService()

	// setup handler
	campaignDetailHandler := handler.NewCampaignDetailsHandler(campaignDetailService, authService)

	// RUN SERVICE
	router := gin.Default()

	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group api
	api := router.Group("api/v1")

	// routing
	api.POST("/create_campaign", middleware.AuthApiCampaignMiddleware(authService, campaignDetailService), campaignDetailHandler.CreateCampaign)
	// update information umkm
	api.PUT("/update_information_umkm", middleware.AuthApiCampaignMiddleware(authService, campaignDetailService), campaignDetailHandler.UpdateInformationUmkm)
	// update proposal umkm file gcp

	// update more information umkm

	// Create Campign With Verify Token User Campaign : Cheack If user (user_campaign_id) have campaign or not, If not create new campaign

	// Create Image Campaign With Verify Token User Campaign and unix_id : Cheack If campign not ready or not, If not create new image campaign

	// Update Image Campaign With Verify Token User Campaign and unix_id : Cheack If campign not ready or not, If not create new image campaign

	// Get Image Campaign With unix_id

	// Update Campaign With Verify Token User Campaign

	// Get Campaign Detail With Verify Token User Campaign

	// // Get All Campaign With Verify Token User Campaign

	// Anonim Accsess
	// Get All Campaign Active : Done Status

	// Get Campaign Detail ? Active : Done Status

	// Admin Accsess
	// Get Campaign Detail With Verify Token User Admin

	// Delete Campaign With Verify Token User Admin

	// Get All Campaign With Verify Token User Admin ? Active : Deactive : Report : Done Status

	// Active Campaign With Verify Token User Admin

	// Deactive Campaign With Verify Token User Admin

	// Admin Service Accsess
	// GET STATUS SERVICE WITH VERIFY TOKEN ADMIN SERVICE

	// GET LOG SERVICE WITH VERIFY TOKEN ADMIN SERVICE

	// Reviewer Accsess

	// Get All Campaign With Verify Token User Reviewer ? Active : Deactive : Report : Done Status

	// Get Campaign Detail With Verify Token User Reviewer ? Active : Deactive : Report Status : Done Status

	// Report Campaign With Verify Token User Reviewer

	// Active Campaign With Verify Token User Reviewer check if campaign have report or not, if not active campaign

	// Investor Accsess
	// Get All Campaign Active With Verify Token User Investor ? Active : Done Status

	// Get Campaign Detail With Verify Token User Investor ? Active : Done Status

	// Invest Campaign With Verify Token User Investor {With Transaction Service}

	// end Rounting
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}

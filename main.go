package main

import (
	"./config"
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "JWT"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 24 * time.Hour,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	bank := router.Group("/bank")
	{
		bank.GET("/", inDB.GetBanks)
		bank.GET("/:id", inDB.GetBankByID)

		bank.POST("/", inDB.CreateBank)
		bank.PUT("/:id", inDB.UpdateBankByID)
		bank.DELETE("/:id", inDB.DeleteBankByID)
	}

	bankAccount := router.Group("/bank-account")
	{
		bankAccount.GET("/", inDB.GetBankAccounts)
		bankAccount.GET("/:id", inDB.GetBankAccountByID)

		bankAccount.POST("/", inDB.CreateBankAccount)
		bankAccount.PUT("/:id", inDB.UpdateBankAccountByID)
		bankAccount.DELETE("/:id", inDB.DeleteBankAccountByID)
	}

	donor := router.Group("/donor")
	{
		donor.GET("/", inDB.GetDonors)
		donor.GET("/:id", inDB.GetDonorByID)
		donor.GET("/:id/article-id", inDB.GetDonorsByArticleID)

		donor.POST("/", inDB.CreateDonor)
		donor.PUT("/:id", inDB.UpdateDonorByID)
		donor.DELETE("/:id", inDB.DeleteDonorByID)
	}

	article := router.Group("/article")
	{
		article.GET("/", inDB.GetArticles)
		article.GET("/:id", inDB.GetArticleByID)

		article.POST("/", inDB.CreateArticle)
		article.PUT("/:id", inDB.UpdateArticleByID)
		article.DELETE("/:id", inDB.DeleteArticleByID)
	}

	request := router.Group("/request")
	{
		request.GET("/", inDB.GetRequests)
		request.GET("/:id", inDB.GetRequestByID)

		request.POST("/", inDB.CreateRequest)
		request.PUT("/:id", inDB.UpdateRequestByID)
		request.DELETE("/:id", inDB.DeleteRequestByID)
	}

	articleImage := router.Group("/image/article")
	{
		articleImage.GET("/", inDB.GetArticleImages)
		articleImage.GET("/:id", inDB.GetArticleImageByID)

		articleImage.POST("/", inDB.CreateArticleImage)
		articleImage.PUT("/:id", inDB.UpdateArticleImageByID)
		articleImage.DELETE("/:id", inDB.DeleteArticleImageByID)
	}

	requestImage := router.Group("/image/request")
	{
		requestImage.GET("/", inDB.GetRequestImages)
		requestImage.GET("/:id", inDB.GetRequestImageByID)

		requestImage.POST("/", inDB.CreateRequestImage)
		requestImage.PUT("/:id", inDB.UpdateRequestImageByID)
		requestImage.DELETE("/:id", inDB.DeleteRequestImageByID)
	}

	user := router.Group("/user")
	{
		user.GET("/", inDB.GetUsers)
		user.GET("/:id", inDB.GetUserByID)

		user.POST("/", inDB.CreateUser)
		user.PUT("/:id", inDB.UpdateUserByID)
		user.DELETE("/:id", inDB.DeleteUserByID)
	}

	auth := router.Group("/auth")
	{
		auth.GET("/:username/:password", inDB.AuthorizeUser)
	}

	err := router.Run(":8117")
	if err != nil {
		panic(err)
	}
}

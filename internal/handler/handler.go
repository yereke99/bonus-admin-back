package handler

import (
	"net/http"

	"bonus-admin-back/config"
	"bonus-admin-back/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service   *service.Services
	zapLogger *zap.Logger
	appConfig *config.Config
}

func NewHandler(service *service.Services, zapLogger *zap.Logger, appConfig *config.Config) *Handler {
	return &Handler{
		service:   service,
		zapLogger: zapLogger,
		appConfig: appConfig,
	}
}

func (h *Handler) InitHandler() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer", "Authority"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return origin == "https://api.worldbonussystem.com" },
	}))

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// Customer Email code send
	r.POST("/request-otp")
	r.POST("/login")

	// Refresh token
	r.POST("/refresh-token")

	// С токеном Admin
	r.GET("/super-admin/companies")

	r.GET("/super-admin/business-centers")
	r.GET("/super-admin/companies/:company_id/company-assets")
	r.POST("/super-admin/companies")        // create company
	r.POST("/super-admin/business-centers") // create business centers

	r.POST("/super-admin/companies/:company_id/company-assets")
	r.PATCH("/super-admin/companies/:id/bonus/increase") // Increase bonus to Company
	r.GET("/customer/transactions")                      // Get Transactions
	r.PATCH("/super-admin/companies/bonus/reset")        // Эмитент
	r.PATCH("/super-admin/company-assets/{assetId}")     // Update asset

	return r
}

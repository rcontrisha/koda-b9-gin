package router

import (
	"rcontrisha/koda-b9-gin/internal/controller"
	"rcontrisha/koda-b9-gin/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService)

	authRouter.POST("/login", authController.Login)
	authRouter.POST("/register", authController.Register)
}

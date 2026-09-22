package router

import (
	"rcontrisha/koda-b9-gin/internal/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authController := controller.NewAuthController()

	authRouter.POST("/login", authController.Login)
	authRouter.POST("/register", authController.Register)
}

package controller

import (
	"fmt"
	"rcontrisha/koda-b9-gin/internal/dto"
	"rcontrisha/koda-b9-gin/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Login(ctx *gin.Context) {
	var payload dto.User

	if e := ctx.ShouldBindWith(&payload, binding.JSON); e != nil {
		return
	}

	if res := service.NewAuthService().LoginService(payload); res != nil {
		ctx.JSON(404, gin.H{
			"status": "failed",
			"msg":    fmt.Sprintln("Login failed. Invalid username or password."),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": "success",
		"msg":    fmt.Sprintf("Login success. Welcome, %s.", payload.Username),
	})
}

func (a *AuthController) Register(ctx *gin.Context) {
	var payload dto.User
	if e := ctx.ShouldBindWith(&payload, binding.JSON); e != nil {
		return
	}

	if res := service.NewAuthService().RegisterService(payload); res != nil {
		ctx.JSON(404, gin.H{
			"status": "failed",
			"msg":    res.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "success",
		"msg":    "Registration success.",
	})
}

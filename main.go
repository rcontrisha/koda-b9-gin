package main

import (
	"fmt"
	"log"

	// "log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Tes struct {
	Tes int
}

type User struct {
	Username string
	Password string
}

var Users = []User{}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "world",
		})
	})

	router.GET("/koda", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "kdoa",
		})
	})

	router.POST("/tes", func(ctx *gin.Context) {
		var data Tes
		if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
			ctx.JSON(500, gin.H{
				"msg": "Error",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"msg": fmt.Sprintf("Tes %d", data.Tes),
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var payload User
		if e := ctx.ShouldBindWith(&payload, binding.JSON); e != nil {
			return
		}
		// log.Println(payload)
		// log.Println(Users)
		if len(Users) > 0 {
			for _, user := range Users {
				if payload.Username == user.Username && payload.Password == user.Password {
					ctx.JSON(200, gin.H{
						"status": "success",
						"msg":    fmt.Sprintf("Login success. Welcome, %s.", payload.Username),
					})
					return
				}
			}

			ctx.JSON(404, gin.H{
				"status": "failed",
				"msg":    fmt.Sprintln("Login failed. Invalid username or password."),
			})
		}
	})

	router.POST("/register", func(ctx *gin.Context) {
		var payload User
		log.Println(Users)
		if e := ctx.ShouldBindWith(&payload, binding.JSON); e != nil {
			return
		}

		if len(payload.Password) < 8 {
			ctx.JSON(400, gin.H{
				"status": "failed",
				"msg":    fmt.Sprintln("Register failed. Password must at least 8 characters."),
			})
			return
		}

		if len(Users) > 0 {
			for _, user := range Users {
				if payload.Username == user.Username {
					ctx.JSON(400, gin.H{
						"status": "failed",
						"msg":    "Register failed. User already exist.",
					})
					return
				}
			}
		}

		Users = append(Users, payload)
		ctx.JSON(200, gin.H{
			"status": "success",
			"msg":    "Registration success.",
		})
	})

	router.Run()
}

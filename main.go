package main

import (
	"rcontrisha/koda-b9-gin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.MainRouter(r)

	r.Run()
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/employee", func(ctx *gin.Context) {
		ctx.File("gin-1-fundamentals-course-PS/routing-with-http-methods/public/employee.html")
	})

	router.POST("/employee", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "New request POSTED successfully")
	})

	log.Fatal(router.Run(":3000"))

}

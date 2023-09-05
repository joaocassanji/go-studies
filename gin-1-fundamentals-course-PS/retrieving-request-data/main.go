package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/query/*rest", func(ctx *gin.Context) {
		username := ctx.Query("username")
		year := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := ctx.QueryArray("month")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	r.GET("/employee", func(ctx *gin.Context) {
		ctx.File("gin-1-fundamentals-course-PS/retrieving-request-data/public/employee.html")
	})

	r.POST("/employee", func(ctx *gin.Context) {
		date := ctx.PostForm("date")
		amount := ctx.PostForm("amount")
		username := ctx.DefaultPostForm("username", "admin")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		})
	})

	log.Fatal(r.Run(":3000"))
}

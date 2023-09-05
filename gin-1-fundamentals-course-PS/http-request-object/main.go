package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/*rest", func(ctx *gin.Context) {
		url := ctx.Request.URL.String()
		headers := ctx.Request.Header
		cookies := ctx.Request.Cookies()

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"headers": headers,
			"cookies": cookies,
		})
	})

	log.Fatal(r.Run(":3000"))
}

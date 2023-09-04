package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	adminGroup := r.Group("/admin")
	adminGroup.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to admininister users")
	})
	adminGroup.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to admininister roles")
	})
	adminGroup.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to admininister policies")
	})

	log.Fatal(r.Run(":3000"))
}

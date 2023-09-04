package main

import (
	"embed"
	"log"

	"github.com/gin-gonic/gin"
)

var f embed.FS

func main() {
	r, _ := createRouter("/", "gin-1-fundamentals-course-PS/serving-files/public/index.html")
	log.Fatal(r.Run(":3000"))

}

func createRouter(routePath, filePath string) (*gin.Engine, gin.IRoutes) {
	router := gin.Default()
	route := router.StaticFile(routePath, filePath)
	return router, route
}

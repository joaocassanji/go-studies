package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	filepath := "gin-1-fundamentals-course-PS/routing-with-http-methods/public/employee.html"
	createHttpRoute(r, "/employee", "GET", true, filepath)
	createHttpRoute(r, "/employee", "POST", false, "New request POSTED successfully")

	log.Fatal(r.Run(":3000"))

}

func createHttpRoute(router *gin.Engine, routePath string, method string, isFile bool, filePathOrStr string) gin.IRoutes {
	var route gin.IRoutes

	handler := func(router *gin.Context) {
		if isFile {
			router.File(filePathOrStr)
		} else {
			router.String(http.StatusOK, filePathOrStr)
		}
	}

	switch method {
	case "POST":
		route = router.POST(routePath, handler)
	case "GET":
		route = router.GET(routePath, handler)
	}

	return route
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeOffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"-" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"-"`
}

func main() {

	r := gin.Default()

	r.GET("/employee", func(ctx *gin.Context) {
		ctx.File("gin-1-fundamentals-course-PS/retrieving-request-data/public/employee.html")
	})

	r.POST("/employee", func(ctx *gin.Context) {
		var timeOffRequest TimeOffRequest
		if err := ctx.ShouldBind(&timeOffRequest); err == nil {
			ctx.JSON(http.StatusOK, timeOffRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	apiGroup := r.Group("/api")
	apiGroup.POST("/timeoff", func(ctx *gin.Context) {
		var timeOffRequest TimeOffRequest
		if err := ctx.ShouldBindJSON(&timeOffRequest); err == nil {
			ctx.JSON(http.StatusOK, timeOffRequest)
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	log.Fatal(r.Run(":3000"))
}

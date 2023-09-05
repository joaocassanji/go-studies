package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type TimeOffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"required,future" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

var ValidatorFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

func main() {

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("future", ValidatorFuture)
	}

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

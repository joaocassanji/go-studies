package main

import (
	"demo/gin-1-fundamentals-course-PS/rendering-html/employee"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("gin-1-fundamentals-course-PS/creating-a-custom-middleware/templates/*")

	//r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	registerRoutes(r)

	r.Run()

}

func registerRoutes(r *gin.Engine) {

	tryToGetEmployee := func(c *gin.Context, employeeIDRaw string) (*employee.Employee, bool) {
		employeeID, err := strconv.Atoi(employeeIDRaw)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return nil, false
		}
		emp, err := employee.Get(employeeID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return nil, false
		}
		return emp, true
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/employees")
	})

	r.GET("/employees", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", employee.GetAll())
	})

	r.GET("/employees/:employeeID", func(c *gin.Context) {
		employeeIDRaw := c.Param("employeeID")

		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			c.HTML(http.StatusOK, "employee.tmpl", *emp)
		}
	})

	r.POST("/employees/:employeeID", func(c *gin.Context) {
		var timeoff employee.TimeOff
		err := c.ShouldBind(&timeoff)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested

		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			c.Redirect(http.StatusFound, "/employees/"+employeeIDRaw)
		}
	})

	g := r.Group("/api/employees", Benchmark)
	{
		g.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, employee.GetAll())
		})
		g.GET("/:employeeID", func(c *gin.Context) {
			employeeIDRaw := c.Param("employeeID")
			if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
				c.JSON(http.StatusOK, *emp)
			}
		})
		g.POST("/:employeeID", func(c *gin.Context) {
			var timeoff employee.TimeOff
			err := c.ShouldBind(&timeoff)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			timeoff.Type = employee.TimeoffTypePTO
			timeoff.Status = employee.TimeoffStatusRequested

			employeeIDRaw := c.Param("employeeID")
			if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
				emp.TimeOff = append(emp.TimeOff, timeoff)
				c.JSON(http.StatusOK, *emp)
			}
		})
	}

	r.Static("/public", "./public")
}

var Benchmark gin.HandlerFunc = func(c *gin.Context) {
	t := time.Now()
	c.Next()
	elapsed := time.Since(t)
	log.Print("Time to process: ", elapsed)
}

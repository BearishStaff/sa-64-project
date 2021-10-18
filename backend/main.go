package main

import (
	"github.com/BearishStaff/sa-64-example/controller"
	"github.com/BearishStaff/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// CheckOut Routes
	r.GET("/checkout", controller.ListCheckOut)
	r.GET("/checkout/:id", controller.GetCheckOut)
	r.POST("/checkout", controller.CreateCheckOut)
	r.PATCH("/checkout", controller.UpdateCheckOut)
	r.DELETE("/checkout/:id", controller.DeleteCheckOut)

	// CheckOut Routes
	r.GET("/checkin", controller.ListCheckIn)
	r.GET("/checkin/:id", controller.GetCheckIn)
	r.POST("/checkin", controller.CreateCheckIn)
	r.PATCH("/checkin", controller.UpdateCheckIn)
	r.DELETE("/checkin/:id", controller.DeleteCheckIn)

	// Customer Routes
	r.GET("/customer", controller.ListCustomer)
	r.GET("/customer/:id", controller.GetCustomer)
	r.POST("/customer", controller.CreateCustomer)
	r.PATCH("/customer", controller.UpdateCustomer)
	r.DELETE("/customer/:id", controller.DeleteCustomer)

	// Customer Routes
	r.GET("/employee", controller.ListEmployee)
	r.GET("/employee/:id", controller.GetEmployee)
	r.POST("/employee", controller.CreateEmployee)
	r.PATCH("/employee", controller.UpdateEmployee)
	r.DELETE("/employee/:id", controller.DeleteEmployee)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()

	}

}

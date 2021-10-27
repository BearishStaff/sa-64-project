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
	r.GET("/check_outs", controller.ListCheckOut)
	r.GET("/check_outs/:id", controller.GetCheckOut)
	r.POST("/check_outs", controller.CreateCheckOut)
	r.PATCH("/check_outs", controller.UpdateCheckOut)
	r.DELETE("/check_outs/:id", controller.DeleteCheckOut)

	// CheckOut Routes
	r.GET("/check_ins", controller.ListCheckIn)
	r.GET("/check_ins/:id", controller.GetCheckIn)
	r.POST("/check_ins", controller.CreateCheckIn)
	r.PATCH("/check_ins", controller.UpdateCheckIn)
	r.DELETE("/check_ins/:id", controller.DeleteCheckIn)

	// Customer Routes
	r.GET("/customers", controller.ListCustomer)
	r.GET("/customers/:id", controller.GetCustomer)
	r.POST("/customers", controller.CreateCustomer)
	r.PATCH("/customers", controller.UpdateCustomer)
	r.DELETE("/customers/:id", controller.DeleteCustomer)

	// Employee Routes
	r.GET("/employees", controller.ListEmployee)
	r.GET("/employees/:id", controller.GetEmployee)
	r.POST("/employees", controller.CreateEmployee)
	r.PATCH("/employees", controller.UpdateEmployee)
	r.DELETE("/employees/:id", controller.DeleteEmployee)

	// Room Routes
	r.GET("/rooms", controller.ListRoom)
	r.GET("/rooms/:id", controller.GetRoom)
	r.POST("/rooms", controller.CreateRoom)
	r.PATCH("/rooms", controller.UpdateRoom)
	r.DELETE("/rooms/:id", controller.DeleteRoom)

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

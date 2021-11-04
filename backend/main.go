package main

import (
	"github.com/BearishStaff/sa-64-example/controller"
	"github.com/BearishStaff/sa-64-example/entity"
	"github.com/BearishStaff/sa-64-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// CheckOut Routes
			protected.GET("/check_outs", controller.ListCheckOut)
			protected.GET("/check_outs/:id", controller.GetCheckOut)
			protected.POST("/check_outs", controller.CreateCheckOut)
			protected.PATCH("/check_outs", controller.UpdateCheckOut)
			protected.DELETE("/check_outs/:id", controller.DeleteCheckOut)

			// CheckIn Routes
			protected.GET("/check_ins", controller.ListCheckIns)
			protected.GET("/check_ins/room", controller.ListCheckInRoom)
			protected.GET("/check_ins/:id", controller.GetCheckIn)
			protected.POST("/check_ins", controller.CreateCheckIn)
			protected.PATCH("/check_ins", controller.UpdateCheckIn)
			protected.DELETE("/check_ins/:id", controller.DeleteCheckIn)

			// Customer Routes
			protected.GET("/customers", controller.ListCustomers)
			protected.GET("/customers/:id", controller.GetCustomer)
			protected.POST("/customers", controller.CreateCustomer)
			protected.PATCH("/customers", controller.UpdateCustomer)
			protected.DELETE("/customers/:id", controller.DeleteCustomer)

			// Employee Routes
			protected.GET("/employees", controller.ListEmployees)
			protected.GET("/employees/:id", controller.GetEmployee)
			protected.POST("/employees", controller.CreateEmployee)
			protected.PATCH("/employees", controller.UpdateEmployee)
			protected.DELETE("/employees/:id", controller.DeleteEmployee)

			// Room Routes
			protected.GET("/rooms", controller.ListRooms)
			protected.GET("/rooms/:id", controller.GetRoom)
			protected.POST("/rooms", controller.CreateRoom)
			protected.PATCH("/rooms", controller.UpdateRoom)
			protected.DELETE("/rooms/:id", controller.DeleteRoom)
		}
	}

	r.POST("/login", controller.Login)
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

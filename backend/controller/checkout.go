package controller

import (
	"net/http"

	"github.com/BearishStaff/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /CheckOut
func CreateCheckOut(c *gin.Context) {

	var employee entity.Employee
	var checkin entity.CheckIn
	var checkout entity.CheckOut
	var customer entity.Customer

	if err := c.ShouldBindJSON(&checkout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// หา checkin ด้วย id
	if tx := entity.DB().Where("id = ?", checkout.CheckInID).First(&checkin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkin not found"})
		return
	}

	// หา employee ด้วย id
	if tx := entity.DB().Where("id = ?", checkout.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", checkout.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 12: สร้าง CheckOut
	co := entity.CheckOut{
		CheckIn:      checkin,  // โยงความสัมพันธ์กับ Entity CheckIn
		Customer:     customer, // โยงความสัมพันธ์กับ Entity Customer
		Employee:     employee, // โยงความสัมพันธ์กับ Entity Employee
		CheckOutTime: checkout.CheckOutTime,
		Condition:    checkout.Condition,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&co).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": co})
}

// GET /check_outs/:id

func GetCheckOut(c *gin.Context) {
	var checkout entity.CheckOut
	id := c.Param("id")
	if err := entity.DB().Preload("CheckIn").Preload("Employee").Preload("Customer").Preload("CheckIn.Room").Raw("SELECT * FROM check_outs WHERE id = ?", id).Scan(&checkout).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkout})
}

// GET /checkout

func ListCheckOut(c *gin.Context) {
	var checkouts []entity.CheckOut
	if err := entity.DB().Preload("CheckIn").Preload("Employee").Preload("Customer").Preload("CheckIn.Room").Raw("SELECT * FROM check_outs").Find(&checkouts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkouts})
}

// DELETE /checkouts/:id

func DeleteCheckOut(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM check_outs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this checkout not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /checkout

func UpdateCheckOut(c *gin.Context) {
	var checkout entity.CheckOut
	if err := c.ShouldBindJSON(&checkout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", checkout.ID).First(&checkout); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkout not found"})
		return
	}

	if err := entity.DB().Save(&checkout).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkout})

}

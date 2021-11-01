package controller

import (
	"net/http"

	"github.com/BearishStaff/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /CheckIn
func CreateCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// GET /checkin/:id
func GetCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	id := c.Param("id")
	if err := entity.DB().Preload("Room").Preload("Employee").Preload("Customer").Raw("SELECT * FROM check_ins WHERE reserve_id = ?", id).Find(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// GET /checkin
func ListCheckIn(c *gin.Context) {
	var checkin []entity.CheckIn
	if err := entity.DB().Preload("Room").Preload("Employee").Preload("Customer").Raw("SELECT * FROM check_ins").Find(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// GET /checkin/room:
func ListCheckInRoom(c *gin.Context) {
	var checkin []entity.CheckIn
	if err := entity.DB().Preload("Room").Preload("Employee").Preload("Customer").Raw("SELECT ci.id, ci.room_id, ci.customer_id, ci.employee_id, ci.date_time FROM check_ins ci LEFT JOIN check_outs co ON ci.id = co.check_in_id  WHERE co.check_in_id is NULL").Find(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// DELETE /checkin/:id
func DeleteCheckIn(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM check_ins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /checkin
func UpdateCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", checkin.ID).First(&checkin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkin not found"})
		return
	}

	if err := entity.DB().Save(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

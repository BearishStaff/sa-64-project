package controller

import (
	"net/http"

	"github.com/BearishStaff/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /Room
func CreateRoom(c *gin.Context) {

	var room entity.Room
	var employee entity.Employee
	var roomtype entity.RoomType
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา roomtype ด้วย id
	if tx := entity.DB().Where("id = ?", room.TypeID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomtype not found"})
		return
	}

	// 9: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", room.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 10: ค้นหา employe ด้วย id
	if tx := entity.DB().Where("id = ?", room.RecorderID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: สร้าง Room
	r := entity.Room{
		Roomnumber: room.Roomnumber,
		Location:   room.Location,
		Type:       roomtype, // โยงความสัมพันธ์กับ Entity RoomType
		Status:     status,   // โยงความสัมพันธ์กับ Entity Status
		Recorder:   employee, // โยงความสัมพันธ์กับ Entity Employee
	}

	// 12: บันทึก
	if err := entity.DB().Create(&r).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": r})
}

// GET /Room/:id
func GetRoom(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	if err := entity.DB().Preload("Type").Preload("Status").Preload("Recorder").Raw("SELECT * FROM rooms WHERE id = ?", id).Find(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}

// GET /Room
func ListRooms(c *gin.Context) {
	var rooms []entity.Room
	if err := entity.DB().Preload("Type").Preload("Status").Preload("Recorder").Raw("SELECT * FROM rooms").Find(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// Delete /rooms/:id
func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /rooms
func UpdateRoom(c *gin.Context) {
	var room entity.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", room.ID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	if err := entity.DB().Save(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

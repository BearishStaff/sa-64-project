package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema

	database.AutoMigrate(
		&CheckOut{},
		&CheckIn{},
		&Customer{},
		&Employee{},
		&Room{},
	)
	db = database

	// ===== สมมติ Employee =====
	password, err := bcrypt.GenerateFromPassword([]byte("111111"), 14)

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Waree Aaaa",
		Tel:      "099-987-6543",
		Email:    "aaa@gmail.com",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Phupha Bbbb",
		Tel:      "02-453-3333",
		Email:    "bbb@gmail.com",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Napha Cccc",
		Tel:      "088-888-8888",
		Email:    "ccc@gmail.com",
		Password: string(password),
	})

	var waree Employee
	var phupha Employee
	var napha Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "aaa@gmail.com").Scan(&waree)
	db.Raw("SELECT * FROM employees WHERE email = ?", "bbb@gmail.com").Scan(&phupha)
	db.Raw("SELECT * FROM employees WHERE email = ?", "ccc@gmail.com").Scan(&napha)

	// ===== สมมติ Customer =====

	db.Model(&Customer{}).Create(&Customer{
		Name:     "Golden Dddd",
		Tel:      "02-222-2222",
		Email:    "ddd@hotmail.com",
		Password: string(password),
	})
	db.Model(&Customer{}).Create(&Customer{
		Name:     "Shepherd Eeee",
		Tel:      "01-111-1111",
		Email:    "eee@hotmail.com",
		Password: string(password),
	})
	var golden Customer
	var shepherd Customer
	db.Raw("SELECT * FROM customers WHERE email = ?", "ddd@hotmail.com").Scan(&golden)
	db.Raw("SELECT * FROM customers WHERE email = ?", "eee@hotmail.com").Scan(&shepherd)

	// ===== สมมติ Room =====
	db.Model(&Room{}).Create(&Room{
		Location:   "3th",
		Roomnumber: "301",
	})
	db.Model(&Room{}).Create(&Room{
		Location:   "3th",
		Roomnumber: "302",
	})
	db.Model(&Room{}).Create(&Room{
		Location:   "4th",
		Roomnumber: "401",
	})
	db.Model(&Room{}).Create(&Room{
		Location:   "4th",
		Roomnumber: "402",
	})

	var room1 Room
	var room2 Room
	var room3 Room
	var room4 Room
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "401").Scan(&room1)
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "402").Scan(&room2)
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "301").Scan(&room3)
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "302").Scan(&room4)

	// ===== สมมติ CheckIn =====
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  golden, // Customer object
		Employee:  phupha,
		Room:      room2, // Room Object
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  shepherd, // Customer object
		Employee:  phupha,
		Room:      room1, // Room Object
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  shepherd, // Customer object
		Employee:  phupha,
		Room:      room3, // Room Object
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  golden, // Customer object
		Employee:  phupha,
		Room:      room3, // Room Object
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  golden, // Customer object
		Employee:  phupha,
		Room:      room4, // Room Object
	})

	var check_in1 CheckIn
	var check_in2 CheckIn
	db.Raw("SELECT * FROM check_ins WHERE id = ?", 1).Scan(&check_in1)
	db.Raw("SELECT * FROM check_ins WHERE id = ?", 2).Scan(&check_in2)

	db.Model(&CheckOut{}).Create(&CheckOut{
		CheckOutTime: time.Now(),
		Customer:     golden,
		Employee:     phupha,
		CheckIn:      check_in1,
		Condition:    "No damage",
	})
	db.Model(&CheckOut{}).Create(&CheckOut{
		CheckOutTime: time.Now(),
		Customer:     golden,
		Employee:     phupha,
		CheckIn:      check_in2,
		Condition:    "No damage",
	})

}

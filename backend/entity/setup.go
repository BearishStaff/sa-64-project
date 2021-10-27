package entity

import (
	"time"

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
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Waree Aaaa",
		Tel:      "099-987-6543",
		Email:    "aaa@gmail.com",
		Password: "111111",
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Phupha Bbbb",
		Tel:      "02-453-3333",
		Email:    "bbb@gmail.com",
		Password: "111111",
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Napha Cccc",
		Tel:      "088-888-8888",
		Email:    "ccc@gmail.com",
		Password: "111111",
	})

	var waree Employee
	var phupha Employee
	var napha Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "aaa@gmail.com").Scan(&waree)
	db.Raw("SELECT * FROM employees WHERE email = ?", "bbb@gmail.com").Scan(&phupha)
	db.Raw("SELECT * FROM employees WHERE email = ?", "ccc@gmail.com").Scan(&napha)

	// ===== สมมติ Customer =====

	db.Model(&Customer{}).Create(&Customer{
		Name:  "Golden Dddd",
		Tel:   "02-222-2222",
		Email: "ddd@hotmail.com",
	})
	db.Model(&Customer{}).Create(&Customer{
		Name:  "Shepherd Eeee",
		Tel:   "01-111-1111",
		Email: "eee@hotmail.com",
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
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "401").Scan(&room1)
	db.Raw("SELECT * FROM rooms WHERE roomnumber = ?", "402").Scan(&room2)

	// ===== สมมติ CheckIn =====
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		CheckIn:   golden, // Customer object
		Employee:  phupha,
		Reserve:   room2, // Room Object
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		CheckIn:   shepherd, // Customer object
		Employee:  phupha,
		Reserve:   room1, // Room Object
	})

	var check_in1 CheckIn
	var check_in2 CheckIn
	db.Raw("SELECT * FROM check_ins WHERE check_in_id = ?", 1).Scan(&check_in1)
	db.Raw("SELECT * FROM check_ins WHERE check_in_id = ?", 2).Scan(&check_in2)

	// ===== สมมติ CheckOut =====
	// db.Model(&CheckOut{}).Create(&CheckOut{
	// 	CheckIn:      check_in1,
	// 	Customer:     golden,
	// 	Employee:     napha,
	// 	CheckOutTime: time.Now(),
	// })
	// db.Model(&CheckOut{}).Create(&CheckOut{
	// 	CheckIn:      check_in2,
	// 	Customer:     golden,
	// 	Employee:     waree,
	// 	CheckOutTime: time.Now(),
	// })
}

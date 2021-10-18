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
	)
	db = database

	// ===== สมมติ Employee =====
	db.Model(&Employee{}).Create(&Employee{
		Name:  "Samoyed",
		Tel:   "099-987-6543",
		Email: "samoly@gmail.com",
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:  "Husky",
		Tel:   "02-453-3333",
		Email: "siby@gmail.com",
	})

	var samoyed Employee
	var husky Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "samoly@gmail.com").Scan(&samoyed)
	db.Raw("SELECT * FROM employees WHERE email = ?", "siby@example.com").Scan(&husky)

	// ===== สมมติ Customer =====

	db.Model(&Customer{}).Create(&Customer{
		Name:  "golden",
		Tel:   "02-222-2222",
		Email: "retriever@hotmail.com",
	})
	db.Model(&Customer{}).Create(&Customer{
		Name:  "shepherd",
		Tel:   "01-111-1111",
		Email: "police_shep@hotmail.com",
	})
	var golden Customer
	var shepherd Customer
	db.Raw("SELECT * FROM customers WHERE email = ?", "retriever@hotmail.com").Scan(&golden)
	db.Raw("SELECT * FROM customers WHERE email = ?", "police_shep@hotmail.com").Scan(&shepherd)

	// ===== สมมติ CheckIn =====
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  golden,
		Employee:  husky,
		Room:      "402",
	})
	db.Model(&CheckIn{}).Create(&CheckIn{
		Date_time: time.Now(),
		Customer:  shepherd,
		Employee:  husky,
		Room:      "401",
	})

	var check_in1 CheckIn
	var check_in2 CheckIn
	db.Raw("SELECT * FROM check_ins WHERE customer_id = ?", 1).Scan(&check_in1)
	db.Raw("SELECT * FROM check_ins WHERE customer_id = ?", 2).Scan(&check_in2)

	// ===== สมมติ CheckOut =====
	db.Model(&CheckOut{}).Create(&CheckOut{
		CheckIn:      check_in1,
		Customer:     golden,
		Employee:     samoyed,
		CheckOutTime: time.Now(),
	})
	db.Model(&CheckOut{}).Create(&CheckOut{
		CheckIn:      check_in2,
		Customer:     golden,
		Employee:     samoyed,
		CheckOutTime: time.Now(),
	})
}

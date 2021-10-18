package entity

import (
	"time"

	"gorm.io/gorm"
)

type CheckOut struct {
	gorm.Model

	CheckInID *uint
	CheckIn   CheckIn

	CustomerID *uint
	Customer   Customer

	EmployeeID *uint
	Employee   Employee

	CheckOutTime time.Time
}

type CheckIn struct {
	gorm.Model
	Date_time time.Time

	CheckInID *uint
	CheckIn   Customer

	Reserve Room

	PaymentID *uint

	EmployeeID *uint
	Employee   Employee
}

type Customer struct {
	gorm.Model
	Name  string
	Tel   string
	Email string

	// 1 customer มีได้หลาย checkOut
	CheckOuts []CheckOut `gorm:"foreignKey:CustomerID"`
}

type Employee struct {
	gorm.Model
	Name  string
	Tel   string
	Email string

	CheckIns []CheckIn `gorm:"foreignKey:EmployeeID"`
	// 1 employee checkout ได้หลายครั้ง
	CheckOuts []CheckOut `gorm:"foreignKey:EmployeeID"`
}

type Room struct {
	gorm.Model
	Location   string
	Roomnumber string
	Records    []CheckIn `gorm:"foreignKey:ReserveID"`
}

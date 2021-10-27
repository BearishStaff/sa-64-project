package entity

import (
	"time"

	"gorm.io/gorm"
)

type CheckOut struct {
	gorm.Model

	CheckInID *uint   `gorm:"unique"`
	CheckIn   CheckIn `gorm:"references:id"`

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	CheckOutTime time.Time
}

type CheckIn struct {
	gorm.Model
	Date_time time.Time

	CheckInID *uint
	CheckIn   Customer `gorm:"references:id"`

	ReserveID *uint
	Reserve   Room `gorm:"references:id"`

	PaymentID *uint

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
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

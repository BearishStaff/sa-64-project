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

	Condition string
}

type CheckIn struct {
	gorm.Model
	Date_time time.Time

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	RoomID *uint
	Room   Room `gorm:"references:id"`

	PaymentID *uint

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	CheckOut []CheckOut `gorm:"foreignKey:CheckInID"`
}

type Customer struct {
	gorm.Model
	Name     string
	Tel      string
	Email    string `gorm:"unique"`
	Password string

	// 1 customer มีได้หลาย checkOut
	CheckOuts []CheckOut `gorm:"foreignKey:CustomerID"`
	CheckIns  []CheckIn  `gorm:"foreignKey:CustomerID"`
}

type Employee struct {
	gorm.Model
	Name     string
	Tel      string
	Email    string `gorm:"unique"`
	Password string

	CheckIns []CheckIn `gorm:"foreignKey:EmployeeID"`
	// 1 employee checkout ได้หลายครั้ง
	CheckOuts []CheckOut `gorm:"foreignKey:EmployeeID"`
}

type Room struct {
	gorm.Model
	Location   string
	Roomnumber string
	Records    []CheckIn `gorm:"foreignKey:RoomID"`
}

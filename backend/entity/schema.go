package entity

import (
	"time"

	"gorm.io/gorm"
)

type CheckIn struct {
	gorm.Model
	DateTime time.Time

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	RoomID *uint
	Room   Room `gorm:"references:id"`

	RoomPaymentID *uint       `gorm:"uniqueIndex"`
	RoomPayment   RoomPayment `gorm:"references:id"`

	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	// เป็นข้อมูล employee เมื่อ join ตาราง
	Employee Employee `gorm:"references:id"`

	CheckOut           []CheckOut          `gorm:"foreignKey:CheckInID"`
	RepairInformations []RepairInformation `gorm:"foreignKey:CheckInID"`
}
type Customer struct {
	gorm.Model
	Name         string
	Email        string `gorm:"uniqueIndex"`
	Password     string
	Tel          string
	CheckIns     []CheckIn     `gorm:"foreignKey:CustomerID"`
	Record       []CheckIn     `gorm:"foreignKey:CustomerID"`
	Reservations []Reservation `gorm:"foreignKey:CustomerID"`
	CheckOuts    []CheckOut    `gorm:"foreignKey:CustomerID"`
}
type Employee struct {
	gorm.Model
	Name         string
	Tel          string
	Email        string `gorm:"uniqueIndex"`
	Password     string
	Records      []CheckIn     `gorm:"foreignKey:EmployeeID"`
	RoomPayments []RoomPayment `gorm:"foreignKey:RecorderID"`
	Rooms        []Room        `gorm:"foreignKey:RecorderID"`
	CheckIns     []CheckIn     `gorm:"foreignKey:EmployeeID"`
	CheckOuts    []CheckOut    `gorm:"foreignKey:EmployeeID"`
}
type RoomPayment struct {
	gorm.Model
	PaymentDate time.Time
	Amount      int
	//Record CheckIn `gorm:"references:id"`
	//RecordID *uint
	//Record []CheckIn `gorm:"foreignKey:PaymentID"`
	RecorderID *uint
	Recorder   Employee

	ReservationID *uint
	Reservation   Reservation

	BalanceID *uint
	Balance   Balance
}
type Room struct {
	gorm.Model
	Location     string
	Roomnumber   string        `gorm:"uniqueIndex"`
	CheckIn      []CheckIn     `gorm:"foreignKey:RoomID"`
	Records      []CheckIn     `gorm:"foreignKey:RoomID"`
	Reservations []Reservation `gorm:"foreignKey:RoomID"`

	TypeID *uint
	Type   RoomType `gorm:"references:id"`
	//EmployeeID ทำหน้าที่เป็น FK
	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`
	//StstusID ทำหน้าที่เป็น FK
	StatusID *uint
	Status   Status `gorm:"references:id"`
}

type CheckOut struct {
	gorm.Model
	CheckInID *uint   `gorm:"uniqueIndex"`
	CheckIn   CheckIn `gorm:"references:id"` // added references
	// CheckIn	[]CheckIn `gorm:"foreignKey:CheckOutID"`
	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	CheckOutTime time.Time

	Condition string
}

type Equipment struct {
	gorm.Model
	Name               string
	RepairInformations []RepairInformation `gorm:"foreignKey:EquipmentID"`
}

type Problem struct {
	gorm.Model
	Value              string
	RepairInformations []RepairInformation `gorm:"foreignKey:ProblemID"`
}

type Urgency struct {
	gorm.Model
	Value              string
	RepairInformations []RepairInformation `gorm:"foreignKey:UrgencyID"`
}

type RepairInformation struct {
	gorm.Model
	Datetime time.Time
	// FK จาก CkeckIn
	CheckInID *uint
	CheckIn   CheckIn
	// FK จาก Equipment
	EquipmentID *uint
	Equipment   Equipment
	// FK จาก Problem
	ProblemID *uint
	Problem   Problem
	// FK จาก Urgency
	UrgencyID *uint
	Urgency   Urgency
}
type Payment struct {
	gorm.Model
	Method       string
	Reservations []Reservation `gorm:"foreignKey:PaymentID"`
}
type Reservation struct {
	gorm.Model
	People      int
	DateAndTime time.Time

	CustomerID *uint
	Customer   Customer

	RoomID *uint
	Room   Room

	PaymentID *uint
	Payment   Payment

	RoomPayments []RoomPayment `gorm:"foreignKey:ReservationID"`
}
type Balance struct {
	gorm.Model

	Type string

	RoomPayments []RoomPayment `gorm:"foreignKey:BalanceID"`
}
type RoomType struct {
	gorm.Model
	Name   string
	Detail string
	Price  int
	Rooms  []Room `gorm:"foreignKey:TypeID"`
}
type Status struct {
	gorm.Model
	Detail string
	Rooms  []Room `gorm:"foreignKey:StatusID"`
}

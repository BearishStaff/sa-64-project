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

	database, err := gorm.Open(sqlite.Open("sapro.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&CheckIn{},
		&Customer{},
		&Employee{},
		&RoomPayment{},
		&Room{},
		&CheckOut{},
		&Payment{},
		&Reservation{},
		&Equipment{},
		&Problem{},
		&Urgency{},
		&RepairInformation{},
		&Balance{},
		&RoomType{},
		&Status{},
	)

	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//employee
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นางสาวพร มณีวรรณ",
		Email:    "porn@gmail.com",
		Tel:      "0883322456",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายสม จันทร์เพ็ญ",
		Email:    "som@gmail.com",
		Tel:      "0885548900",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นางสาวกล้วย ไชยวาที",
		Email:    "naruemon@gmail.com",
		Tel:      "0610091572",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Phupha Bbbb",
		Email:    "bbb@gmail.com",
		Tel:      "0945333333",
		Password: string(password), // password change
	})

	var porn Employee
	var som Employee
	var kluy Employee
	var Phupha Employee

	db.Raw("SELECT * FROM employees WHERE email = ?", "porn@gmail.com").Scan(&porn)
	db.Raw("SELECT * FROM employees WHERE email = ?", "som@gmail.com").Scan(&som)
	db.Raw("SELECT * FROM employees WHERE email = ?", "naruemon@gmail.com").Scan(&kluy)
	db.Raw("SELECT * FROM employees WHERE email = ?", "bbb@gmail.com").Scan(&Phupha)

	// Customer Data
	C1 := Customer{
		Name:     "นายแซม  พูลสวัสดิ์",
		Email:    "b6100124@g.sut.ac.th",
		Tel:      "0983322403",
		Password: string(password),
	}
	db.Model(&Customer{}).Create(&C1)

	C2 := Customer{
		Name:     "นายเบนซ์ ปักโคทานัง",
		Email:    "Tanapol@gmail.com",
		Tel:      "0673322403",
		Password: string(password),
	}
	db.Model(&Customer{}).Create(&C2)

	C3 := Customer{
		Name:     "นายหยกชาย พงศธร",
		Email:    "name@example.com",
		Tel:      "0953322883",
		Password: string(password),
	}
	db.Model(&Customer{}).Create(&C3)

	var sam Customer
	var benz Customer
	var yok Customer
	db.Raw("SELECT * FROM customers WHERE email = ?", "b6100124@g.sut.ac.th").Scan(&sam)
	db.Raw("SELECT * FROM customers WHERE email = ?", "Tanapol@gmail.com").Scan(&benz)
	db.Raw("SELECT * FROM customers WHERE email = ?", "name@example.com").Scan(&yok)

	//roomtype
	Standard := RoomType{
		Name:   "standard",
		Detail: "เตียงคู่ 1 เตียง เข้าพักได้ 2 คน มีสิ่งอำนวยความสะดวก (เครื่องปรับอากาศ, น้ำอุ่น)",
		Price:  500,
	}
	db.Model(&RoomType{}).Create(&Standard)

	Deluxe := RoomType{
		Name:   "Deluxe",
		Detail: "เตียงขนาดกลาง 1 เตียง มองเห็นวิวทิวทัศนียภาพที่สวยงาม มีสิ่งอำนวยความสะดวก (เครื่องปรับอากาศ, น้ำอุ่น)",
		Price:  1000,
	}
	db.Model(&RoomType{}).Create(&Deluxe)

	Suite := RoomType{
		Name:   "suite",
		Detail: "เตียงขนาดใหญ่  1 เตียง มีห้องนั่งเล่น ห้องนอน มีสิ่งอำนวยความสะดวก (เครื่องปรับอากาศ, น้ำอุ่น, โซฟา, ตู้เย็น, โทรทัศน์)",
		Price:  1500,
	}
	db.Model(&RoomType{}).Create(&Suite)

	//Room Data
	r1 := Room{
		Roomnumber: "101",
		Location:   "ZoneA ชั้น 1",
	}
	db.Model(&Room{}).Create(&r1)

	r2 := Room{
		Roomnumber: "201",
		Location:   "ZoneB ชั้น 2",
	}
	db.Model(&Room{}).Create(&r2)

	r3 := Room{
		Roomnumber: "301",
		Location:   "ZoneA ชั้น 3",
	}
	db.Model(&Room{}).Create(&r3)

	// Payment
	p1 := Payment{
		Method: "KTB",
	}
	db.Model(&Payment{}).Create(&p1)

	p2 := Payment{
		Method: "SCB",
	}
	db.Model(&Payment{}).Create(&p2)

	p3 := Payment{
		Method: "TMB",
	}
	db.Model(&Payment{}).Create(&p3)

	//reserve
	//reserve 1
	RS1 := Reservation{
		DateAndTime: time.Now(),
		People:      1,
		Customer:    sam,
		Room:        r1,
		Payment:     p1,
	}
	//reserve 2
	RS2 := Reservation{
		DateAndTime: time.Now(),
		People:      2,
		Customer:    yok,
		Room:        r2,
		Payment:     p3,
	}
	//reserve 3
	RS3 := Reservation{
		DateAndTime: time.Now(),
		People:      1,
		Customer:    benz,
		Room:        r2,
		Payment:     p1,
	}

	db.Model(&Reservation{}).Create(&RS1)
	db.Model(&Reservation{}).Create(&RS2)
	db.Model(&Reservation{}).Create(&RS3)

	//balance
	B1 := Balance{
		Type: "Full",
	}
	db.Model(&Balance{}).Create(&B1)

	B2 := Balance{
		Type: "Half",
	}
	db.Model(&Balance{}).Create(&B2)

	// --- RoomPayment Data
	RP1 := RoomPayment{
		PaymentDate: time.Now(),
		Amount:      600,
		Recorder:    som,
		Reservation: RS1,
		Balance:     B2,
	}
	db.Model(&RoomPayment{}).Create(&RP1)

	RP2 := RoomPayment{
		PaymentDate: time.Now(),
		Amount:      1200,
	}
	db.Model(&RoomPayment{}).Create(&RP2)

	RP3 := RoomPayment{
		PaymentDate: time.Now(),
		Amount:      2400,
	}
	db.Model(&RoomPayment{}).Create(&RP3)

	// ---  CheckIn data
	CheckIntana1 := CheckIn{
		DateTime:    time.Now(),
		Customer:    benz,
		Room:        r1,
		RoomPayment: RP2,
		Employee:    kluy,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana1)

	CheckIntana2 := CheckIn{
		DateTime:    time.Now(),
		Customer:    sam,
		Room:        r2,
		RoomPayment: RP1,
		Employee:    som,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana2)

	CheckIntana3 := CheckIn{
		DateTime:    time.Now(),
		Customer:    yok,
		Room:        r3,
		RoomPayment: RP3,
		Employee:    porn,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana3)

	// checkout data
	CheckOuttana3 := CheckOut{
		CheckIn:      CheckIntana3,
		Employee:     som,
		Customer:     yok,
		CheckOutTime: time.Time{},
		Condition:    "No damage",
	}
	db.Model(&CheckOut{}).Create(&CheckOuttana3)

	// equipment data

	equipdm := Equipment{
		Name: "กระจกสำหรับแต่งตัว (Dressing mirror)",
	}
	db.Model(&Equipment{}).Create(&equipdm)

	equipchair := Equipment{
		Name: "เก้าอี้ (Chair)",
	}
	db.Model(&Equipment{}).Create(&equipchair)

	equipwaterheater := Equipment{
		Name: "เครื่องทำน้ำอุ่น (Water heater)",
	}
	db.Model(&Equipment{}).Create(&equipwaterheater)

	equiplamp := Equipment{
		Name: "โคมไฟ (Lamp)",
	}
	db.Model(&Equipment{}).Create(&equiplamp)

	equipflushtoilet := Equipment{
		Name: "ชักโครก (flush toilet)",
	}
	db.Model(&Equipment{}).Create(&equipflushtoilet)

	equipbed := Equipment{
		Name: "เตียงนอน (Bed)",
	}
	db.Model(&Equipment{}).Create(&equipbed)

	equipfridge := Equipment{
		Name: "ตู้เย็น (Fridge)",
	}
	db.Model(&Equipment{}).Create(&equipfridge)

	equiptable := Equipment{
		Name: "โต๊ะ (Table)",
	}
	db.Model(&Equipment{}).Create(&equiptable)

	equipwardrobe := Equipment{
		Name: "ตู้เสื้อผ้า (Wardrobe)",
	}
	db.Model(&Equipment{}).Create(&equipwardrobe)

	equiptv := Equipment{
		Name: "ทีวี (TV)",
	}
	db.Model(&Equipment{}).Create(&equiptv)

	equipdoor := Equipment{
		Name: "ประตู (Door)",
	}
	db.Model(&Equipment{}).Create(&equipdoor)

	equipshower := Equipment{
		Name: "ฝักบัว (Shower)",
	}
	db.Model(&Equipment{}).Create(&equipshower)

	equipfan := Equipment{
		Name: "พัดลม (Fan)",
	}
	db.Model(&Equipment{}).Create(&equipfan)

	equipFluorescentlamp := Equipment{
		Name: "หลอดฟลูออเรสเซนต์ (Fluorescent lamp)",
	}
	db.Model(&Equipment{}).Create(&equipFluorescentlamp)

	equipac := Equipment{
		Name: "แอร์ (Air conditioner)",
	}
	db.Model(&Equipment{}).Create(&equipac)

	// problem data
	prodefective := Problem{
		Value: "ชำรุด (Defective)",
	}
	db.Model(&Problem{}).Create(&prodefective)

	pronotwork := Problem{
		Value: "ใช้งานไม่ได้ (Not working)",
	}
	db.Model(&Problem{}).Create(&pronotwork)

	// urgency data
	urgent := Urgency{
		Value: "เร่งด่วน (Urgent)",
	}
	db.Model(&Urgency{}).Create(&urgent)

	urfast := Urgency{
		Value: "เร็ว (Fast)",
	}
	db.Model(&Urgency{}).Create(&urfast)

	urmedium := Urgency{
		Value: "ปานกลาง (Medium)",
	}
	db.Model(&Urgency{}).Create(&urmedium)

	urslow := Urgency{
		Value: "ช้า (Slow)",
	}
	db.Model(&Urgency{}).Create(&urslow)

	urvslow := Urgency{
		Value: "ช้ามากๆ (Vary slow)",
	}
	db.Model(&Urgency{}).Create(&urvslow)

	// 1 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana1,
		Equipment: equipfridge,
		Problem:   pronotwork,
		Urgency:   urfast,
		Datetime:  time.Now(),
	})

	// 2 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana2,
		Equipment: equiptable,
		Problem:   prodefective,
		Urgency:   urmedium,
		Datetime:  time.Now(),
	})

	// 3 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana3,
		Equipment: equipwardrobe,
		Problem:   prodefective,
		Urgency:   urslow,
		Datetime:  time.Now(),
	})

	//status
	available := Status{
		Detail: "ว่าง",
	}
	db.Model(&Status{}).Create(&available)

	noavailable := Status{
		Detail: "ไม่ว่าง",
	}
	db.Model(&Status{}).Create(&noavailable)

	//
	// === Query
	//

	// var target Employee
	// db.Model(&Employee{}).Find(&target, db.Where("email = ?", "chanwit@gmail.com"))

	// var employeeLogin Employee
	// db.Model(&Employee{}).Find(&employeeLogin, db.Where("email = ? and employee_id = ?", "พนักงานที่เข้าสู่ระบบ", target.ID))

	/*var watchedList []*WatchVideo
	db.Model(&WatchVideo{}).
		Joins("Playlist").
		Joins("Resolution").
		Joins("Video").
		Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	for _, wl := range watchedList {
		fmt.Printf("Watch Video: %v\n", wl.ID)
		fmt.Printf("%v\n", wl.Playlist.Title)
		fmt.Printf("%v\n", wl.Resolution.Value)
		fmt.Printf("%v\n", wl.Video.Name)
		fmt.Println("====")
	}*/

}

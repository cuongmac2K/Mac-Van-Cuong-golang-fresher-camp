package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Deparment struct {
	MaPB   int    `json:"MaPB,omitempty" gorm:"column:MaPB;"`
	HoTen  string `json:"HoTen" gorm:"column:HoTen;"`
	DiaChi string `json:"DiaChi" gorm:"column:DiaChi;"`
	SDT    int    `json:"SDT" gorm:"column:SDT;"`
}

type DeparmentUpdate struct {
	MaPB   int    `json:"MaPB,omitempty" gorm:"column:MaPB;"`
	HoTen  string `json:"HoTen" gorm:"column:HoTen;"`
	DiaChi string `json:"DiaChi" gorm:"column:DiaChi;"`
	SDT    int    `json:"SDT" gorm:"column:SDT;"`
}

func (Deparment) TableName() string {
	return "department"
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/customer?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("k ket noi dc ", err)
	}

	// Insert data restaurant
	newData := Deparment{MaPB: 3, HoTen: "Nhan Su", DiaChi: "Hai Duong", SDT: 888123456}

	if err := db.Create(&newData); err != nil {
		fmt.Println(err)
	}
	fmt.Println(newData)

	// find all resaults
	var departments []Deparment
	db.Where("HoTen=? ", "Ke toan").Find(&departments)
	db.Where("address=? ", "okla").Find(&departments)
	fmt.Println(departments)

	//tim kiem theo ma PB
	var department Deparment
	if err := db.Where("MaPB=? ", 1).First(&department); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(departments)
	}
	// Update

	department.DiaChi = "HCM"
	db.Table(Deparment{}.TableName()).Where("MaPB = 1").Updates(&department)
	fmt.Println(departments)

	///Update with nil value
	DiaChi1 := "okla"
	db.Table(Deparment{}.TableName()).Where("MaPB = 1").Updates(&DeparmentUpdate{DiaChi: DiaChi1})
	fmt.Println(departments)

	//Delete
	db.Table(Deparment{}.TableName()).Where("MaPB = 0").Delete(nil)
	fmt.Println(departments)
}

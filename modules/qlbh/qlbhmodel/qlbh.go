package qlbhmodel

type Nhanvien struct {
	Id       int    `json:"id" gorm:"column:id;"`
	Name     string `json:"Name" gorm:"column:Name;"`
	Age      int    `json:"Age" gorm:"column:Age;"`
	Adress   string `json:"Adress" gorm:"column:Adress;"`
	ThamNien int    `json:"ThamNien" gorm:"column:ThamNien;"`
}

func (Nhanvien) TableName() string {
	return "nhanvien"
}

type NhanvienUpdate struct {
	Name   *string `json:"Name" gorm:"column:Name;"`
	Adress *string `json:"Adress" gorm:"column:Adress;"`
}

func (NhanvienUpdate) TableName() string {
	return Nhanvien{}.TableName()
}

type QlbhCreate struct {
	Id       int    `json:"id" gorm:"column:id;"`
	Name     string `json:"Name" gorm:"column:Name;"`
	Age      int    `json:"Age" gorm:"column:Age;"`
	Adress   string `json:"Adress" gorm:"column:Adress;"`
	ThamNien int    `json:"ThamNien" gorm:"column:ThamNien;"`
}

func (QlbhCreate) TableName() string {
	return Nhanvien{}.TableName()
}

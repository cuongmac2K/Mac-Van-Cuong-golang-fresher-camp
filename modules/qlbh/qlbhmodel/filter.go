package qlbhmodel

type Filter struct {
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth;"`
}

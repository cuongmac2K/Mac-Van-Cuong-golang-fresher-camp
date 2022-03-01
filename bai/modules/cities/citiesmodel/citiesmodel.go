package citiesmodel

import (
	"demo/common"
	"errors"
	"strings"
)

const EntityName = "Cities"

type Cities struct {
	common.SQLModel `json:"inline"`
	Title           string `json:"title" gorm:"title;"`
	Status          int    `json:"status" gorm:"status"`
}

func (Cities) TableName() string {
	return "cities"
}

type CitiesUpdate struct {
	Title  string `json:"title" gorm:"title;"`
	Status int    `json:"status" gorm:"status"`
}

func (CitiesUpdate) TableName() string {
	return Cities{}.TableName()
}

type CitiesCreate struct {
	Id     int    `json:"id" gorm:"colum:id;"`
	Title  string `json:"title"gorm:"title";`
	Status int    `json:"status"gorm:"status";`
}

func (res *CitiesCreate) TableName() string {
	return Cities{}.TableName()
}
func (res *CitiesCreate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)
	if len(res.Title) == 0 {
		return errors.New("NO empty")
	}
	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "cities name can't be blank", "ErrNameCannotBeEmpty")
)

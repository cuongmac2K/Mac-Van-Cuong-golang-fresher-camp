package restaurantlikemodel

import (
	"demo/common"
	"errors"
	"fmt"
	"time"
)

const EntityName = "UserlikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm :"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("can not like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrUserLikedRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("you have already liked restaurant"),
		fmt.Sprintf("ErrUserLikedRestaurant"),
	)
}

func ErrCannotUnLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("can not unlike restaurant"),
		fmt.Sprintf("ErrCannotUnLikeRestaurant"),
	)
}

var (
	ErrUserNotLikeRestaurant = common.NewCustomError(
		errors.New("you have not liked restaurant"),
		"you have not liked restaurant",
		"ErrUserNotLikeRestaurant",
	)

	ErrUserLikeRestaurant = common.NewCustomError(
		errors.New("you have already liked restaurant"),
		"you have already liked restaurant",
		"ErrUserLikeRestaurant",
	)
)

package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"restaurant_id,omitempty" form:"restaurant_id"`
	UserId       int `json:"user_id "gorm:"user_id"`
}

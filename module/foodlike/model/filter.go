package foodlikemodel

type Filter struct {
	FoodId int `json:"-" form:"food_id"`
	UserId int `json:"-" form:"user_id"`
}

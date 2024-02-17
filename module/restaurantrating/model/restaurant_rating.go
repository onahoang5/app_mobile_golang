package restaurantratingmodel

import "Food-delivery/common"

const EntityName = "UserRatingRestaurant"

type Rating struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int                `json:"-" gorm:"column:restaurant_id;"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	Point           float64            `json:"point" gorm:"column:point;"`
	Comment         string             `json:"comment" gorm:"column:comment;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Rating) TableName() string {
	return "restaurant_ratings"
}

func (r *Rating) GetRestaurantId() int {
	return r.RestaurantId
}

func (r *Rating) GetUserId() int {
	return r.UserId
}

func (r *Rating) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type CreateRatingRestaurant struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int     `json:"-" gorm:"column:restaurant_id;"`
	UserId          int     `json:"-" gorm:"column:user_id;"`
	Point           float64 `json:"point" gorm:"column:point;"`
	Comment         string  `json:"comment" gorm:"column:comment;"`
}

func (CreateRatingRestaurant) TableName() string {
	return Rating{}.TableName()
}

type UpdateRatingRestaurant struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int `json:"-" gorm:"column:restaurant_id;"`
	UserId          int `json:"-" gorm:"column:user_id;"`
	// Point           float64 `json:"point" gorm:"column:point;"`
	Comment string `json:"comment" gorm:"column:comment;"`
}

func (UpdateRatingRestaurant) TableName() string {
	return Rating{}.TableName()
}

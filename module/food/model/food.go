package foodmodel

import (
	"Food-delivery/common"
	"errors"
	"fmt"
)

type FoodType string

const EntityName = "Food"

type Food struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int                      `json:"-" gorm:"column:restaurant_id;"`
	CategoryId      int                      `json:"-" gorm:"column:category_id;"`
	Name            string                   `json:"name" gorm:"column:name;"`
	Type            FoodType                 `json:"type" gorm:"column:type;"`
	Description     string                   `json:"description" gorm:"column:description;"`
	Price           float64                  `json:"price" gorm:"column:price;"`
	Images          *common.Image            `json:"images" gorm:"column:images;"`
	LikedCount      int                      `json:"liked_count" gorm:"column:liked_count;"` // computed field
	UserId          int                      `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser       `json:"user" gorm:"Preload:false;"`
	Restaurant      *common.SimpleRestaurant `json:"restaurant" gorm:"Preload:false;"`
}

func (Food) TableName() string {
	return "foods"
}

func (data *Food) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeFood)

	if u := data.Restaurant; u != nil {
		u.Mask(isAdminOrOwner)
	}

}

// func (f *Food) GetUserId() int {
// 	return f.UserId
// }

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"-" gorm:"column:restaurant_id;"`
	CategoryId      int           `json:"-" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Price           int           `json:"price" gorm:"price;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}

type FoodUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Price       int           `json:"price" gorm:"price;"`
	Images      *common.Image `json:"images" gorm:"column:images;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

func ErrAlreadyFoodRestaurant() *common.AppError {
	return common.NewCustomError(
		errors.New("Already create food this restaurant"),
		fmt.Sprintf("Already create food this restaurant"),
		fmt.Sprintf("ERR_ALREADY_FOOD_RESTAURANT"),
	)
}

func ErrCannotCreateFoodRestaurant() *common.AppError {
	return common.NewCustomError(
		errors.New("Already create food this restaurant"),
		fmt.Sprintf("Already create food this restaurant"),
		fmt.Sprintf("ERR_ALREADY_FOOD_RESTAURANT"),
	)
}

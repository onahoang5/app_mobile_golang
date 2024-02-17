package restaurantmodel

import (
	"Food-delivery/common"
	"errors"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremiun RestaurantType = "premium"
const EntityName = "Restaurant"

var (
	ErrNameCannotEmpty = errors.New("restaurant name cannot be blank")
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	// Id              int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
	// Status int            `json:"status" gorm:"column:status"`
	Type       RestaurantType     `json:"type" gorm:"column:type;"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	LikedCount int                `json:"liked_count" gorm:"column:liked_count;"` // computed field
	UserId     int                `json:"-" gorm:"column:user_id;"`
	User       *common.SimpleUser `json:"user" gorm:"Preload:false;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) GetUserId() int {
	return r.UserId
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	// Id   int    `json:"id" gorm:"column:id;"`
	Name   string         `json:"name" gorm:"column:name;"`
	Addr   string         `json:"addr" gorm:"column:addr;"`
	UserId int            `json:"-" gorm:"column:user_id;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover  *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmtpy
	}
	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string        `json:"name" gorm:"column:name;"`
	Addr            *string        `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmtpy = errors.New("name cannot be emtpy")
)

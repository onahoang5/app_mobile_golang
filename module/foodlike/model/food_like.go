package foodlikemodel

import (
	"Food-delivery/common"
	"errors"
	"fmt"

	"time"
)

const EntityName = "UserFoods"

type Likefood struct {
	FoodId    int                `json:"food_id" gorm:"column:food_id;"`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Likefood) TableName() string {
	return "food_likes"
}

func (l *Likefood) GetFoodId() int {
	return l.FoodId
}

func (l *Likefood) GetUserId() int {
	return l.UserId
}

func ErrCannotLikeFood(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this Food"),
		fmt.Sprintf("ERR_CANNOT_LIKE_FOOD"),
	)
}

func ErrAlreadyLikedFood() *common.AppError {
	return common.NewCustomError(
		errors.New("Already liked this Food"),
		fmt.Sprintf("Already liked this Food"),
		fmt.Sprintf("ERR_ALREADY_LIKED_FOOD"),
	)
}

func ErrAlreadyUnLikedFood() *common.AppError {
	return common.NewCustomError(
		errors.New("Already unliked this Food"),
		fmt.Sprintf("Already unliked this Food"),
		fmt.Sprintf("ERR_ALREADY_UNLIKED_FOOD"),
	)
}

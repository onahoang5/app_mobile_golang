package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	DbTypeFood       = 3
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDisLikeRestaurant = "TopicUserDisLikeRestaurant"
	TopicUserLikeFood          = "TopicUserLikeFood"
	TopicUserDisLikeFood       = "TopicUserDisLikeFood"
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery Error", err)
	}
}

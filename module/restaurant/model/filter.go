package restaurantmodel

type Filter struct {
	OwnerId int   `json:"owner_id,omitemtpy" form:"owner_id"`
	Status  []int `json:"-"`
}

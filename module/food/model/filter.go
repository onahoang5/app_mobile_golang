package foodmodel

type Filter struct {
	CategoryId   int   `json:"category_id,omitempty" form:"category_id"`
	RestaurantId int   `json:"restaurant_id,omitempty" form:"restaurant_id"`
	Status       []int `json:"-"`
}

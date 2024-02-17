package common

type SimpleRestaurant struct {
	SQLModel `json:",inline"`
	Name     string  `json:"name" gorm:"column:name;"`
	Addr     string  `json:"addr" gorm:"column:addr;"`
	Logo     *Image  `json:"logo" gorm:"column:logo;"`
	Cover    *Images `json:"cover" gorm:"column:cover;"`
}

func (SimpleRestaurant) TableName() string {
	return "restaurants"
}

func (u *SimpleRestaurant) Mask(isAdmin bool) {
	u.GenUID(DbTypeRestaurant)
}

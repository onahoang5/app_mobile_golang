package restaurantstorage

import "gorm.io/gorm"

type sqlsStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlsStore {
	return &sqlsStore{db: db}
}

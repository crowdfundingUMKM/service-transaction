package core

import "gorm.io/gorm"

type Repository interface {
}

// "go.mongodb.org/mongo-driver/mongo"
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

package core

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
}

// "go.mongodb.org/mongo-driver/mongo"
type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

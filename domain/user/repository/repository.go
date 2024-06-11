package repository

import (
	"backend-nabati/infrastructure/database"
)

type UserRepository interface {
}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		Database: db,
	}
}

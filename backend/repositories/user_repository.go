package repositories

import (
	"gorm.io/gorm"
	"backend/entities/users"
)

type UserRepository struct {
	*BaseRepository[users.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[users.User](db),
	};
}

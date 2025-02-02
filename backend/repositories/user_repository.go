package repositories

import (
	"gorm.io/gorm"
	"backend/entities/users"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db};
}

func (userRepo *UserRepository) CreateUser(user *users.User) error{
	return userRepo.db.Create(user).Error;
}

func (userRepo *UserRepository) GetUserById(id int64) (*users.User, error) {
	user := &users.User{};
	err := userRepo.db.First(user, id).Error;
	return user, err;
}

func (userRepo *UserRepository) GetUserByEmail(email string) (*users.User, error) {
	user := &users.User{};
	err := userRepo.db.Where("email = ?", email).First(user).Error;
	return user, err;
}

func (userRepo *UserRepository) UpdateUser(user *users.User) error {
	return userRepo.db.Save(user).Error;
}

func (userRepo *UserRepository) DeleteUser(id int64) error {
	return userRepo.db.Delete(&users.User{}, id).Error;
}

func (userRepo *UserRepository) GetAllUsers() ([]users.User, error) {
	var users []users.User
	if err := userRepo.db.Find(&users).Error; err != nil {
		return nil, err;
	}
	return users, nil;
}
package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/satyajitnayk/uptime_monitor/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create inserts a new user into the database
func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by their ID from the database
func (ur *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package database

import (
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var userEntity entity.User
	if err := u.DB.Where("email = ?", email).First(&userEntity).Error; err != nil {
		return nil, err
	}
	return &userEntity, nil
}

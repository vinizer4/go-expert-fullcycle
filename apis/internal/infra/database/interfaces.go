package database

import (
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

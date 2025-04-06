package database

import (
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/entity/user"
)

type UserInterface interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}

package user

import (
	"example.local/go-pilot/internal/models"
)

// Repository represent the user's repository contract
type Repository interface {
	Fetch() (res []*models.User, err error)
	Create(u *models.User) error
	Update(uid int, u *models.User) error
	Delete(uid int) error
}

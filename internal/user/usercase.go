package user

import (
	"example.local/users/internal/models"
)

// Usecase represet user's usecases
type Usecase interface {
	Fetch() (res []*models.User, err error)
	Create(u *models.User) error
	Update(uid int, u *models.User) error
	Delete(uid int) error
}

package usecase

import (
	"example.local/users/internal/models"
	"example.local/users/internal/user"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create new userUsecase object representation of user.Usecase interface
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

// Fetch will extract all users
func (u *userUsecase) Fetch() ([]*models.User, error) {
	users, err := u.userRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create will create new user entity
func (u *userUsecase) Create(m *models.User) error {
	err := u.userRepo.Create(m)
	if err != nil {
		return err
	}

	return nil
}

// Update will update user entity based on uid
func (u *userUsecase) Update(uid int, m *models.User) error {
	err := u.userRepo.Update(uid, m)
	if err != nil {
		return err
	}

	return nil
}

// Delete will remove user entity based on uid
func (u *userUsecase) Delete(uid int) error {
	err := u.userRepo.Delete(uid)
	if err != nil {
		return err
	}

	return nil
}

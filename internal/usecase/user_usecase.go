package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id_user uint) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(id_user uint, user *entity.User) error
	DeleteUser(id_user uint) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		userRepo: repository.NewUserRepository(db),
	}
}
func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.FindAll()
}
func (u *userUsecase) GetUserByID(id uint) (entity.User, error) {
	return u.userRepo.FindByID(id)
}
func (u *userUsecase) GetUserByEmail(email string) (entity.User, error) {
	return u.userRepo.FindByEmail(email)
}
func (u *userUsecase) CreateUser(user *entity.User) error {
	return u.userRepo.Create(user)
}
func (u *userUsecase) UpdateUser(id uint, user *entity.User) error {
	return u.userRepo.Update(id, user)
}
func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}

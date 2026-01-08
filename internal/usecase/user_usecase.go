package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id_user uint) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user *entity.User) error
	RegisterUser(name, email, password string) (*entity.User, error)
	VerifyLogin(email, password string) (*entity.User, error)
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

func (u *userUsecase) RegisterUser(name, email, password string) (*entity.User, error) {
	// Cek apakah email sudah terdaftar
	existingUser, err := u.userRepo.FindByEmail(email)
	if err == nil && existingUser.ID != 0 {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *userUsecase) VerifyLogin(email, password string) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Password == "" {
		return nil, errors.New("please login with Google")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

func (u *userUsecase) UpdateUser(id uint, user *entity.User) error {
	return u.userRepo.Update(id, user)
}
func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}

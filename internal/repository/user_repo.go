package repository

import (
	"api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByID(id uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Create(user *entity.User) error
	Update(id uint, user *entity.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *userRepository) FindByID(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return user, err
}
func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) Update(id uint, user *entity.User) error {
	return r.db.Model(&entity.User{}).
		Where("id_user = ?", id).
		Updates(user).Error
}
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

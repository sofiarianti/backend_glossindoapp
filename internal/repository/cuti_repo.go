package repository

import (
	"api/internal/entity"
	"gorm.io/gorm"
)

type CutiRepository interface {
	FindAll() ([]entity.Cuti, error)
	FindByID(id uint) (entity.Cuti, error)
	FindByUserID(userID uint) ([]entity.Cuti, error)
	Create(cuti *entity.Cuti) error
	Update(id uint, cuti *entity.Cuti) error
	Delete(id uint) error
}

type cutiRepository struct {
	db *gorm.DB
}

func NewCutiRepository(db *gorm.DB) CutiRepository {
	return &cutiRepository{db}
}

func (r *cutiRepository) FindAll() ([]entity.Cuti, error) {
	var cutis []entity.Cuti
	if err := r.db.Find(&cutis).Error; err != nil {
		return nil, err
	}
	return cutis, nil
}

func (r *cutiRepository) FindByID(id uint) (entity.Cuti, error) {
	var cuti entity.Cuti
	err := r.db.First(&cuti, id).Error
	return cuti, err
}

func (r *cutiRepository) FindByUserID(userID uint) ([]entity.Cuti, error) {
	var cutis []entity.Cuti
	err := r.db.Where("id_user = ?", userID).Find(&cutis).Error
	return cutis, err
}

func (r *cutiRepository) Create(cuti *entity.Cuti) error {
	return r.db.Create(cuti).Error
}

func (r *cutiRepository) Update(id uint, cuti *entity.Cuti) error {
	return r.db.Model(&entity.Cuti{}).
		Where("id = ?", id).
		Updates(cuti).Error
}

func (r *cutiRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Cuti{}, id).Error
}

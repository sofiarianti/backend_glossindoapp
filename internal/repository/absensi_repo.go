package repository

import (
	"api/internal/entity"
	"time"

	"gorm.io/gorm"
)

type AbsensiRepository interface {
	FindAll() ([]entity.Absensi, error)
	FindByID(id uint) (entity.Absensi, error)
	FindByUserID(userID uint) ([]entity.Absensi, error)
	FindByUserAndDate(userID uint, date time.Time) (*entity.Absensi, error) // Added this
	Create(absensi *entity.Absensi) error
	Update(id uint, absensi *entity.Absensi) error
	Delete(id uint) error
}

type absensiRepository struct {
	db *gorm.DB
}

func NewAbsensiRepository(db *gorm.DB) AbsensiRepository {
	return &absensiRepository{db}
}

func (r *absensiRepository) FindAll() ([]entity.Absensi, error) {
	var absensis []entity.Absensi
	if err := r.db.Find(&absensis).Error; err != nil {
		return nil, err
	}
	return absensis, nil
}

func (r *absensiRepository) FindByID(id uint) (entity.Absensi, error) {
	var absensi entity.Absensi
	err := r.db.First(&absensi, id).Error
	return absensi, err
}

func (r *absensiRepository) FindByUserID(userID uint) ([]entity.Absensi, error) {
	var absensis []entity.Absensi
	err := r.db.Where("user_id = ?", userID).Find(&absensis).Error
	return absensis, err
}

// FindByUserAndDate finds an attendance record for a specific user on a specific date
func (r *absensiRepository) FindByUserAndDate(userID uint, date time.Time) (*entity.Absensi, error) {
	var absensi entity.Absensi
	// We use DATE() function to compare only the date part
	err := r.db.Where("user_id = ? AND DATE(tanggal) = DATE(?)", userID, date).First(&absensi).Error
	if err != nil {
		return nil, err
	}
	return &absensi, nil
}

func (r *absensiRepository) Create(absensi *entity.Absensi) error {
	return r.db.Create(absensi).Error
}

func (r *absensiRepository) Update(id uint, absensi *entity.Absensi) error {
	return r.db.Model(&entity.Absensi{}).
		Where("id = ?", id).
		Updates(absensi).Error
}

func (r *absensiRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Absensi{}, id).Error
}

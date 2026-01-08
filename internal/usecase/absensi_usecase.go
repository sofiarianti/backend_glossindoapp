package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AbsensiUsecase interface {
	GetAllAbsensis() ([]entity.Absensi, error)
	GetAbsensiByID(id uint) (entity.Absensi, error)
	GetAbsensiByUserID(userID uint) ([]entity.Absensi, error)
	CreateAbsensi(absensi *entity.Absensi) error
	UpdateAbsensi(id uint, absensi *entity.Absensi) error
	DeleteAbsensi(id uint) error
	CheckIn(absensi *entity.Absensi) error
	CheckOut(userID uint, checkoutData *entity.Absensi) error
}

type absensiUsecase struct {
	absensiRepo repository.AbsensiRepository
}

func NewAbsensiUsecase(db *gorm.DB) AbsensiUsecase {
	return &absensiUsecase{
		absensiRepo: repository.NewAbsensiRepository(db),
	}
}

func (u *absensiUsecase) GetAllAbsensis() ([]entity.Absensi, error) {
	return u.absensiRepo.FindAll()
}

func (u *absensiUsecase) GetAbsensiByID(id uint) (entity.Absensi, error) {
	return u.absensiRepo.FindByID(id)
}

func (u *absensiUsecase) GetAbsensiByUserID(userID uint) ([]entity.Absensi, error) {
	return u.absensiRepo.FindByUserID(userID)
}

func (u *absensiUsecase) CreateAbsensi(absensi *entity.Absensi) error {
	return u.absensiRepo.Create(absensi)
}

func (u *absensiUsecase) UpdateAbsensi(id uint, absensi *entity.Absensi) error {
	return u.absensiRepo.Update(id, absensi)
}

func (u *absensiUsecase) DeleteAbsensi(id uint) error {
	return u.absensiRepo.Delete(id)
}

func (u *absensiUsecase) CheckIn(absensi *entity.Absensi) error {
	// 1. Check if already checked in today
	existing, err := u.absensiRepo.FindByUserAndDate(absensi.UserID, time.Now())
	if err == nil && existing != nil {
		return errors.New("already checked in today")
	}

	// 2. Set default values
	absensi.Tanggal = time.Now()
	absensi.CheckInTime = time.Now()
	absensi.Status = "Hadir" // Default status
	
	// Ensure CheckOutTime is nil (handled by pointer in entity or just ignored)
	
	return u.absensiRepo.Create(absensi)
}

func (u *absensiUsecase) CheckOut(userID uint, checkoutData *entity.Absensi) error {
	// 1. Find today's attendance
	existing, err := u.absensiRepo.FindByUserAndDate(userID, time.Now())
	if err != nil {
		return errors.New("no check-in record found for today")
	}

	// 2. Update checkout time
	now := time.Now()
	existing.CheckOutTime = &now
	
	// Update location/distance info if provided
	if checkoutData.Latitude != 0 {
		existing.Latitude = checkoutData.Latitude
	}
	if checkoutData.Longitude != 0 {
		existing.Longitude = checkoutData.Longitude
	}
	if checkoutData.Alamat != "" {
		existing.Alamat = checkoutData.Alamat
	}
	// ... Update other fields if necessary
	
	return u.absensiRepo.Update(existing.ID, existing)
}

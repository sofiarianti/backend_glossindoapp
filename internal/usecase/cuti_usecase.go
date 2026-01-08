package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"gorm.io/gorm"
)

type CutiUsecase interface {
	GetAllCutis() ([]entity.Cuti, error)
	GetCutiByID(id uint) (entity.Cuti, error)
	GetCutiByUserID(userID uint) ([]entity.Cuti, error)
	CreateCuti(cuti *entity.Cuti) error
	UpdateCuti(id uint, cuti *entity.Cuti) error
	DeleteCuti(id uint) error
}

type cutiUsecase struct {
	cutiRepo repository.CutiRepository
}

func NewCutiUsecase(db *gorm.DB) CutiUsecase {
	return &cutiUsecase{
		cutiRepo: repository.NewCutiRepository(db),
	}
}

func (u *cutiUsecase) GetAllCutis() ([]entity.Cuti, error) {
	return u.cutiRepo.FindAll()
}

func (u *cutiUsecase) GetCutiByID(id uint) (entity.Cuti, error) {
	return u.cutiRepo.FindByID(id)
}

func (u *cutiUsecase) GetCutiByUserID(userID uint) ([]entity.Cuti, error) {
	return u.cutiRepo.FindByUserID(userID)
}

func (u *cutiUsecase) CreateCuti(cuti *entity.Cuti) error {
	return u.cutiRepo.Create(cuti)
}

func (u *cutiUsecase) UpdateCuti(id uint, cuti *entity.Cuti) error {
	return u.cutiRepo.Update(id, cuti)
}

func (u *cutiUsecase) DeleteCuti(id uint) error {
	return u.cutiRepo.Delete(id)
}

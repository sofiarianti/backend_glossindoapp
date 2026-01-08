package entity

import (
	"time"
)

type Cuti struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         uint      `gorm:"column:id_user;not null" json:"id_user"`
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai;type:date" json:"tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai;type:date" json:"tanggal_selesai"`
	JenisCuti      string    `gorm:"column:jenis_cuti;type:varchar(100)" json:"jenis_cuti"`
	Status         string    `gorm:"column:status;type:enum('Menunggu','Disetujui','Ditolak');default:'Menunggu'" json:"status"`
	Alasan         string    `gorm:"column:alasan;type:text" json:"alasan"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

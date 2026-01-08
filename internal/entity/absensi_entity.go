package entity

import "time"

type Absensi struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID  uint      `gorm:"column:user_id;not null;uniqueIndex:idx_user_tanggal" json:"user_id"`
	Tanggal time.Time `gorm:"column:tanggal;type:date;not null;uniqueIndex:idx_user_tanggal" json:"tanggal"`
	CheckInTime  time.Time  `gorm:"column:check_in_time" json:"check_in_time"`
	CheckOutTime *time.Time `gorm:"column:check_out_time" json:"check_out_time,omitempty"`
	Latitude  float64 `gorm:"column:latitude;type:decimal(10,8)" json:"latitude"`
	Longitude float64 `gorm:"column:longitude;type:decimal(11,8)" json:"longitude"`
	Alamat    string  `gorm:"column:alamat;type:text" json:"alamat"`
	Distance        float64 `gorm:"column:distance;type:double" json:"distance"`
	IsValidLocation bool    `gorm:"column:is_valid_location" json:"is_valid_location"`
	Status string `gorm:"column:status;type:varchar(20)" json:"status"`
	Source string `gorm:"column:source;type:varchar(20)" json:"source"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

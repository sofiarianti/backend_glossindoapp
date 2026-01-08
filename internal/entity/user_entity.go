package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"`
	GoogleID  string    `gorm:"column:google_id;unique" json:"google_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Email     string    `gorm:"column:email;not null;unique" json:"email"`
	PhotoURL  string    `gorm:"column:photo_url" json:"photo_url"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

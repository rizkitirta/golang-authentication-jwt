package entity

import "time"

type User struct {
	ID        int       `gorm:"primary_key:auto_increment:not null" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	Password  string    `gorm:"->;<-;type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	Token     string    `gorm:"-" json:"token,omitempty"`
}
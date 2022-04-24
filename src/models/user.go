package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int        `gorm:"column:id;" json:"id"`
	Username  string     `gorm:"column:username;" json:"username"`
	Email     string     `gorm:"column:email;" json:"email"`
	Password  string     `gorm:"column:password" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt
}

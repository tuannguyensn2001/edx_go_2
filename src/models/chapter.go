package models

import (
	"gorm.io/gorm"
	"time"
)

type Chapter struct {
	Id        int             `json:"id" gorm:"column:id"`
	Name      string          `json:"name" gorm:"column:name"`
	CourseId  int             `json:"courseId" gorm:"column:course_id"`
	Order     int             `json:"order" gorm:"column:order"`
	CreatedAt *time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
	Course    Course          `json:"course"`
}

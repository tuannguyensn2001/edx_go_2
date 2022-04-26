package models

import (
	enumcourse "edx_go_2/src/enums/course"
	"gorm.io/gorm"
	"time"
)

type Course struct {
	Id          int               `json:"id" gorm:"column:id"`
	Name        string            `json:"name" gorm:"column:name"`
	Description string            `json:"description" gorm:"column:description"`
	ImageUrl    string            `json:"imageUrl" gorm:"column:imageUrl"`
	Price       float64           `json:"price" gorm:"column:price"`
	Status      enumcourse.Status `json:"status" gorm:"column:status"`
	UserId      int               `json:"userId" gorm:"column:user_id"`
	CreatedAt   *time.Time        `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   *time.Time        `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt
	Chapters    []Chapter `json:"chapters"`
}

func (Course) TableName() string {
	return "courses"
}

package courserepository

import (
	"context"
	"edx_go_2/src/models"
	coursestruct "edx_go_2/src/modules/course/struct"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, input coursestruct.CreateCourseInput) (*models.Course, error) {
	course := models.Course{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Status:      input.Status,
		UserId:      input.UserId,
		Price:       input.Price,
	}

	if err := r.db.Create(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil

}

func (r *Repository) FindById(ctx context.Context, id int) (*models.Course, error) {
	var course models.Course

	err := r.db.Where("id = ?", id).Preload("Chapters.Course").First(&course).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *Repository) Update(ctx context.Context, id int, input coursestruct.UpdateCourseInput) (*models.Course, error) {

	course, err := r.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	course.Name = input.Name
	course.Description = input.Description
	course.Price = input.Price
	course.ImageUrl = input.ImageUrl
	course.Status = input.Status

	err = r.db.Save(&course).Error

	if err != nil {
		return nil, err
	}

	return course, nil

}

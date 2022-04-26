package chapterrepository

import (
	"context"
	"edx_go_2/src/models"
	chapterstruct "edx_go_2/src/modules/chapter/struct"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetLatestOrder(ctx context.Context, courseId int) (int64, error) {
	var count int64

	if err := r.db.Model(&models.Chapter{}).Where("course_id = ?", courseId).Count(&count).Error; err != nil {
		return -1, err
	}

	return count, nil
}

func (r *Repository) Create(ctx context.Context, chapter models.Chapter) (*models.Chapter, error) {
	if err := r.db.Create(&chapter).Error; err != nil {
		return nil, err
	}

	return &chapter, nil
}

func (r *Repository) FindById(ctx context.Context, id int) (*models.Chapter, error) {
	var result models.Chapter

	if err := r.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Update(ctx context.Context, id int, input chapterstruct.UpdateChapterInput) (*models.Chapter, error) {
	chapter, err := r.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	chapter.Name = input.Name

	err = r.db.Save(chapter).Error

	if err != nil {
		return nil, err
	}

	return chapter, nil

}

func (r *Repository) Delete(ctx context.Context, id int) error {
	err := r.db.Delete(&models.Chapter{}, id).Error

	return err
}

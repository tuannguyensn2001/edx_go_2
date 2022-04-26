package courseservice

import (
	"context"
	"edx_go_2/src/app"
	"edx_go_2/src/models"
	courserepository "edx_go_2/src/modules/course/repository"
	coursestruct "edx_go_2/src/modules/course/struct"
	"errors"
)

type Service struct {
	repository *courserepository.Repository
}

func NewService(repository *courserepository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, input coursestruct.CreateCourseInput) (*models.Course, error) {
	return s.repository.Create(ctx, input)
}

func (s *Service) Update(ctx context.Context, id int, input coursestruct.UpdateCourseInput, userId int) (*models.Course, error) {
	course, err := s.repository.Update(ctx, id, input)

	if err != nil {
		return nil, err
	}

	if course.UserId != userId {
		return nil, app.ErrNoPermission(errors.New("User don't have permission to update course"))
	}

	return course, nil

}

func (s *Service) FindById(ctx context.Context, id int) (*models.Course, error) {
	course, err := s.repository.FindById(ctx, id)

	if err != nil {
		return nil, app.ErrEntityNotFound(err, "Khong tim thay khoa hoc nay")
	}

	return course, nil

}
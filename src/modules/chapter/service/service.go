package chapterservice

import (
	"context"
	"edx_go_2/src/models"
	chapterrepository "edx_go_2/src/modules/chapter/repository"
	chapterstruct "edx_go_2/src/modules/chapter/struct"
)

type Service struct {
	repository *chapterrepository.Repository
}

func NewService(repository *chapterrepository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, input chapterstruct.CreateChapterInput) (*models.Chapter, error) {
	latestOrder, err := s.repository.GetLatestOrder(ctx, input.CourseId)

	if err != nil {
		return nil, err
	}

	chapter := models.Chapter{
		Name:     input.Name,
		CourseId: input.CourseId,
		Order:    int(latestOrder) + 1,
	}

	result, err := s.repository.Create(ctx, chapter)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *Service) Update(ctx context.Context, id int, input chapterstruct.UpdateChapterInput) (*models.Chapter, error) {
	chapter, err := s.repository.Update(ctx, id, input)

	if err != nil {
		return nil, err
	}

	return chapter, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

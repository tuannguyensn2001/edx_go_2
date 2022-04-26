package authservice

import (
	"context"
	"edx_go_2/src/models"
)

func (s *Service) Me(ctx context.Context, userId int) (*models.User, error) {
	return s.repository.FindById(ctx, userId)
}

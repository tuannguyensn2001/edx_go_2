package authservice

import (
	"context"
	"edx_go_2/src/models"
	authstruct "edx_go_2/src/modules/auth/struct"
)

func (s *Service) Register(ctx context.Context, input authstruct.RegisterInput) (*models.User, error) {
	user, err := s.repository.Create(ctx, input)
	return user, err
}

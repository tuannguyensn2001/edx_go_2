package authservice

import (
	"context"
	"edx_go_2/src/app"
	"edx_go_2/src/models"
	authstruct "edx_go_2/src/modules/auth/struct"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func (s *Service) Register(ctx context.Context, input authstruct.RegisterInput) (*models.User, error) {

	check, err := s.repository.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if check != nil {
		return nil, app.ErrEntityExisted(errors.New("User existed"), "User existed")
	}

	user, err := s.repository.Create(ctx, input)

	if err != nil {
		return nil, app.NewErrorResponse("Register failed", http.StatusInternalServerError, nil, err)
	}

	return user, err

}

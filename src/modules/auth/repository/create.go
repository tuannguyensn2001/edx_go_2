package authrepository

import (
	"context"
	"edx_go_2/src/models"
	authstruct "edx_go_2/src/modules/auth/struct"
	hashpkg "edx_go_2/src/packages/hash"
)

func (r *Repository) Create(ctx context.Context, data authstruct.RegisterInput) (*models.User, error) {

	password, err := hashpkg.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

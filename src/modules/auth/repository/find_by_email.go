package authrepository

import (
	"context"
	"edx_go_2/src/models"
)

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = @email", map[string]interface{}{"email": email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

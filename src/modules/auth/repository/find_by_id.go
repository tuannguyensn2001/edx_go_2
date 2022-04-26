package authrepository

import (
	"context"
	"edx_go_2/src/models"
)

func (r *Repository) FindById(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	query := r.db.Where("id = ?", id).First(&user)

	if err := query.Error; err != nil {
		return nil, err
	}

	return &user, nil

}

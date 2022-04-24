package authservice

import authrepository "edx_go_2/src/modules/auth/repository"

type Service struct {
	repository *authrepository.Repository
}

func NewService(repository *authrepository.Repository) *Service {
	return &Service{repository: repository}
}

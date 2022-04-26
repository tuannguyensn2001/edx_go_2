package authservice

import authrepository "edx_go_2/src/modules/auth/repository"

type Service struct {
	repository *authrepository.Repository
	secretKey  string
}

func NewService(repository *authrepository.Repository, secretKey string) *Service {
	return &Service{repository: repository, secretKey: secretKey}
}

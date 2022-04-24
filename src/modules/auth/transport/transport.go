package authtransport

import (
	"edx_go_2/src/config"
	authrepository "edx_go_2/src/modules/auth/repository"
	authservice "edx_go_2/src/modules/auth/service"
)

type transport struct {
	service *authservice.Service
	conf    config.Config
}

func NewTransport() *transport {

	t := transport{
		conf: config.Conf,
	}

	r := authrepository.NewRepository(t.conf.Db)
	s := authservice.NewService(r)

	t.service = s

	return &t
}

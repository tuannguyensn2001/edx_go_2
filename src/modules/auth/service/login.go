package authservice

import (
	"context"
	"edx_go_2/src/app"
	authstruct "edx_go_2/src/modules/auth/struct"
	hashpkg "edx_go_2/src/packages/hash"
	jwtpkg "edx_go_2/src/packages/jwt"
	"errors"
)

func (s *Service) Login(ctx context.Context, input authstruct.LoginInput) (*authstruct.LoginOutput, error) {

	user, err := s.repository.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	check := hashpkg.Compare(input.Password, user.Password)

	if !check {
		return nil, app.ErrInvalidRequestWithMessage(errors.New("Wrong password"), "Email or password not valid")
	}

	accessToken, err := jwtpkg.GenerateToken(s.secretKey, user.Id, 10000)

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwtpkg.GenerateToken(s.secretKey, user.Id, 20000)

	if err != nil {
		return nil, err
	}

	return &authstruct.LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil

}

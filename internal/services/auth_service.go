package services

import (
	"fmt"

	"github.com/ronaldalds/res/internal/models"
	"github.com/ronaldalds/res/internal/schemas"
)

func (s *Service) Login(req schemas.LoginRequest) (*models.User, error) {
	user, err := s.GormStore.GetUserByUsernameOrEmail(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to login: username or password is incorrect")
	}
	if !user.Active {
		return nil, fmt.Errorf("failed to login: user is inactive")
	}
	return user, nil
}

func (s *Service) SetToken(id uint, access string) error {
	err := s.RedisStore.SetKey(fmt.Sprintf("%d", id), access, 0)
	if err != nil {
		return fmt.Errorf("failed to set key redis: %s", err.Error())
	}
	return nil
}

package service

import (
	"examples/server/config"
	"examples/server/internal/dao"
)

var UserServiceImpl *UserService

type UserService struct {
	UserDao *dao.UserDao
	Config  config.Config
}

func NewUserService(config config.Config) (*UserService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *UserService) CheckPassword(username, password string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *UserService) GetUsernameByApiKey(apikey string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *UserService) GetApiKeyByUsername(username string) string {
	_ = "STUB: not implemented"
	return ""
}

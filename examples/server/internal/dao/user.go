package dao

import (
	"examples/server/config"
	"examples/server/internal/model"
)

const (
	UsersSectionName = ""
	UsersFileName    = "users.ini"
)

type UserDao struct {
	Config config.Config
	fs     *FileStorage
}

func NewUserDao(config config.Config) (*UserDao, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *UserDao) CreateUser(user model.User) error { _ = "STUB: not implemented"; return nil }

// ValidatePassword 验证密码
func (d *UserDao) ValidatePassword(username, password string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *UserDao) Delete(username string) error { _ = "STUB: not implemented"; return nil }

func (d *UserDao) List() []model.User { _ = "STUB: not implemented"; return nil }

package dao

import (
	"examples/server/config"
	"examples/server/internal/model"
)

const (
	SettingsSectionName = ""
	SettingsFileName    = "settings.ini"
)

type UserSettingDao struct {
	Config config.Config
	fs     *FileStorage
}

func NewUserSettingDao(config config.Config, namespace string) (*UserSettingDao, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *UserSettingDao) Save(key, value string) error { _ = "STUB: not implemented"; return nil }

func (d *UserSettingDao) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (d *UserSettingDao) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (d *UserSettingDao) Setting() model.UserSetting {
	_ = "STUB: not implemented"
	return *new(model.UserSetting)
}

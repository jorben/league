package model

import "league/database"

// AutoMigrate 同步模型与schema
func AutoMigrate() error {
	return database.GetInstance().AutoMigrate(
		User{},
		UserSocialInfo{},
		CasbinRule{},
		Menu{},
	)
}

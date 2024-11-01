package migrations

import (
	"OrdentTest/config"
	"OrdentTest/models"
)

func Migrate() {
	// Migrate the schema
	err := config.DB.AutoMigrate(&models.User{}, &models.Comment{}, &models.Article{})
	if err != nil {
		panic(err)
	}
	SeedRoles(config.DB)
}

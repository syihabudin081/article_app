package migrations

import (
	"OrdentTest/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "User"},
		{Name: "Guest"},
		{Name: "Author"},
	}
	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			panic(err) // Handle the error appropriately in a real app
		}
	}
}

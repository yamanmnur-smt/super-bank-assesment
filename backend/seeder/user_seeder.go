package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/services"
	"yamanmnur/simple-dashboard/pkg/db"
)

func UserSeeder() {
	pass, _ := services.HashPassword("password")
	users := []models.User{
		{Name: "Yaman", Username: "yaman", Password: pass},
	}
	gormdb := db.GetDbHandler().DB
	for _, user := range users {
		gormdb.FirstOrCreate(&user, models.User{Username: user.Username})
	}
}

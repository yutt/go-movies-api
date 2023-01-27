package initializers

import "github.com/yutt/go-movies-api/model"

func SyncDB() {
	DB.AutoMigrate(&model.Film{})
	DB.AutoMigrate(&model.User{})
}

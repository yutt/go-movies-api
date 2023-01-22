package initializers

import "github.com/yutt/go-movies-api/model"

func SyncDB() {
	DB.AutoMigrate(&model.User{})
}

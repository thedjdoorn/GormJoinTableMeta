package repositories

import (
	"GormJoinTableMeta/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var _connection *gorm.DB

func connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.SetupJoinTable(&models.Actor{}, "Movies", &models.Role{})
	db.SetupJoinTable(&models.Movie{}, "Actors", &models.Role{})
	db.AutoMigrate(&models.Actor{}, models.Movie{}, models.Role{})

	_connection = db

	return db
}

package repositories

import "GormJoinTableMeta/models"

func AddMovie(m models.Movie) models.Movie {
	connection().Create(&m)
	return m
}

func GetMovies() []models.Movie {
	var result []models.Movie
	connection().Model(&models.Movie{}).Preload("Actors").Find(&result)
	return result
}

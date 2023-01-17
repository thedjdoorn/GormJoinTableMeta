package repositories

import "GormJoinTableMeta/models"

func AddActor(a models.Actor) models.Actor {
	connection().Create(&a)
	return a
}

func GetActors() []models.Actor {
	var result []models.Actor
	connection().Model(&models.Actor{}).Preload("Movies").Find(&result)
	return result
}

package repositories

import "GormJoinTableMeta/models"

func AssociateRole(actor models.Actor, movie models.Movie, roleName string) {
	newRole := models.Role{MovieId: movie.ID, ActorId: actor.ID, RoleName: roleName}
	connection().Create(&newRole)
}

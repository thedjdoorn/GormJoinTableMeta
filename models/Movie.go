package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name   string
	Actors []*Actor `gorm:"many2many:roles"`
	Role   string   `gorm:"-"`
}

func (m *Movie) AfterFind(tx *gorm.DB) error {
	var Roles []Role
	tx.Where("movie_id = ?", m.ID).Find(&Roles)
	var actors []*Actor
	for _, actor := range m.Actors {
		for _, role := range Roles {
			a := *actor
			if actor.ID == role.ActorId {
				a.Role = role.RoleName
				actors = append(actors, &a)
			}
		}
	}
	m.Actors = actors
	return nil
}

package models

import (
	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	Name   string
	Movies []*Movie `gorm:"many2many:roles"`
	Role   string   `gorm:"-"`
}

func (a *Actor) AfterFind(tx *gorm.DB) error {
	var Roles []Role
	tx.Where("actor_id = ?", a.ID).Find(&Roles)
	var movies []*Movie
	for _, movie := range a.Movies {
		for _, role := range Roles {
			m := *movie
			if movie.ID == role.MovieId {
				m.Role = role.RoleName
				movies = append(movies, &m)
			}
		}
	}
	a.Movies = movies
	return nil
}

package main

import (
	"GormJoinTableMeta/models"
	"GormJoinTableMeta/repositories"
	"fmt"
)

func main() {
	daisy := repositories.AddActor(models.Actor{Name: "Daisy Ridley"})
	adam := repositories.AddActor(models.Actor{Name: "Adam Driver"})

	tfa := repositories.AddMovie(models.Movie{Name: "Star Wars: Episode VII - The Force Awakens"})
	moe := repositories.AddMovie(models.Movie{Name: "Murder on the Orient Express"})
	ms := repositories.AddMovie(models.Movie{Name: "Marriage Story"})

	repositories.AssociateRole(daisy, tfa, "Rey")
	repositories.AssociateRole(daisy, moe, "Mary Debenham")
	repositories.AssociateRole(adam, tfa, "Kylo Ren")
	repositories.AssociateRole(adam, ms, "Charlie")

	movies := repositories.GetMovies()

	for _, movie := range movies {
		fmt.Println("Actors in", movie.Name)
		for _, actor := range movie.Actors {
			fmt.Println(fmt.Sprintf("    %s (%s)", actor.Name, actor.Role))
		}
	}

	fmt.Println("---")

	actors := repositories.GetActors()
	for _, actor := range actors {
		fmt.Println("Movies featuring", actor.Name)
		for _, movie := range actor.Movies {
			fmt.Println(fmt.Sprintf("    %s (%s)", movie.Name, movie.Role))
		}
	}
}

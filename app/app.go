package app

import (
	"golangBootcamp/m/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.setRoutes()
}

func (a *App) setRoutes() {
	a.Router.GET("/pokemons", controllers.FindPokemons)
	a.Router.GET("/pokemons/:id", controllers.FindPokemonById)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

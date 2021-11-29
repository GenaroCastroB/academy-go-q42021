package app

import (
	"golangBootcamp/m/common"
	"golangBootcamp/m/controllers"
	"golangBootcamp/m/repositories"
	"golangBootcamp/m/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	Router         *gin.Engine
	pokemonHandler controllers.PokemonServiceHandler
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	setConfigs()
	if err := setConfigs(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	a.injectDependencies()
	a.setRoutes()
}

func (a *App) injectDependencies() {
	pokemonRepo := repositories.NewPokemonRepo(common.NewCsvReader())
	pokemonService := services.NewPokemonService(pokemonRepo)
	a.pokemonHandler = controllers.NewPokemonServiceHandler(pokemonService)
}

func (a *App) setRoutes() {
	a.Router.GET("/pokemons", a.pokemonHandler.FindPokemons)
	a.Router.GET("/pokemons/:id", a.pokemonHandler.FindPokemonById)
	a.Router.PUT("/load/pokemons", a.pokemonHandler.LoadPokemons)
}

func setConfigs() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("./config")

	return viper.ReadInConfig()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

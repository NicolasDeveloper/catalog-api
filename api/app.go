package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com.br/NicolasDeveloper/catalog-api/api/controllers"
	"github.com.br/NicolasDeveloper/catalog-api/api/ioc"

	"github.com.br/NicolasDeveloper/catalog-api/adapters"
	"github.com.br/NicolasDeveloper/catalog-api/api/common"
	"github.com/golobby/container"

	"go.uber.org/dig"

	"github.com/gorilla/mux"
)

//App initialize app
type App struct {
	router    *mux.Router
	container *dig.Container
}

//NewApp contructor
func NewApp() *App {
	return &App{}
}

//Initialize configure app
func (a *App) Initialize() *App {
	ioc.NewContainer().RegisterDependencies()
	return a
}

//StartupDatabase connect to database
func (a *App) StartupDatabase() *App {
	var dbcontext *adapters.DbContext
	container.Make(&dbcontext)
	dbcontext.Connect()
	return a
}

//ConfigEndpoints configure api version and declare api's endpoints
func (a *App) ConfigEndpoints() *App {
	a.router = mux.NewRouter()
	s := a.router.PathPrefix("/catalog-api/v1/").Subrouter()

	for _, b := range initBundles() {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", a.router)
	return a
}

//Run startup api and alocate a port from os enviroment variable 'APP_PORT' or set default ":3201"
func (a *App) Run() *App {
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = ":3201"
	}

	fmt.Printf("Catalog Api Running on Port: %s\n", port)

	log.Fatal(http.ListenAndServe(port, a.router))

	return a
}

func initBundles() []common.Bundle {
	return []common.Bundle{
		controllers.NewProductRouter(),
	}
}

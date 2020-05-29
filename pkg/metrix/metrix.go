package metrix

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents a route attached to the mux
type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

// App represents server with dependencies
type App struct {
	Router *mux.Router
}

// Init configures server with routes and injected dependencies
func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (a *App) initRoutes() {
	// add any core app routes
}

// AddRoutes takes a slice of Route objects to add to the mux
func (a *App) AddRoutes(routes []Route) {
	for _, r := range routes {
		a.Router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}

// Run spins up web server with associated app configuration
func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

package gitlab

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	a.Router.HandleFunc("/hook/gitlab", HookHandler).Methods("POST")
}

// Run spins up web server with associated app configuration
func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

// HookHandler handles webhooks received from GitLab
func HookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

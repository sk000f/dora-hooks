package gitlab

import (
	"net/http"

	"github.com/sk000f/metrix/pkg/metrix"
)

// parse handles webhooks received from GitLab
func parse(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-GitLab-Event")

	if event == "dummy" {
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusBadRequest)
}

// InitRoutes sets up routes for GitLab hooks
func InitRoutes() []metrix.Route {
	r := []metrix.Route{
		{
			Path:    "/hook/gitlab",
			Handler: parse,
			Method:  http.MethodPost,
		},
	}

	return r
}

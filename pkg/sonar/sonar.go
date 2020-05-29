package sonar

import (
	"net/http"

	"github.com/sk000f/metrix/pkg/metrix"
)

// HookHandler handles webhooks received from GitLab
func HookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// InitRoutes sets up routes for GitLab hooks
func InitRoutes() []metrix.Route {
	r := []metrix.Route{
		{
			Path:    "/hook/sonar",
			Handler: HookHandler,
			Method:  http.MethodPost,
		},
	}

	return r
}

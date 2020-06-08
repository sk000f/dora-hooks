package gitlab

import (
	"net/http"

	"github.com/sk000f/metrix/pkg/metrix"
)

// parse handles webhooks received from GitLab
func parse(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get(HookHeader)

	switch event {
	case PipelineEvent:
		w.WriteHeader(http.StatusOK)
	default:
		// log out message the hook event is invalid
		w.WriteHeader(http.StatusBadRequest)
	}
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

// HookHeader is the standard GitLab hook header
const HookHeader = "X-GitLab-Event"

// PipelineEvent is the X-GitLab-Event header for pipeline hooks
const PipelineEvent string = "Pipeline Hook"

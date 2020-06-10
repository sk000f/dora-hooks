package gitlab

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sk000f/metrix/pkg/metrix"
)

// ParseEvent processes the JSON event body and returns event object or an error
func ParseEvent(r *http.Request) (interface{}, error) {

	if r.Body == nil {
		return nil, errors.New("missing event data")
	}

	defer func() {
		_ = r.Body.Close()
	}()

	eventData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// error reading event data
		return nil, err
	}

	if len(eventData) == 0 {
		// error with actual event data
	}

	return eventData, nil
}

// parse handles webhooks received from GitLab
func handleEvent(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get(HookHeader)

	switch event {
	case PipelineEvent:
		_, err := ParseEvent(r)
		fmt.Println(err)
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
			Handler: handleEvent,
			Method:  http.MethodPost,
		},
	}

	return r
}

// HookHeader is the standard GitLab hook header
const HookHeader = "X-Gitlab-Event"

// PipelineEvent is the X-GitLab-Event header for pipeline hooks
const PipelineEvent string = "Pipeline Hook"

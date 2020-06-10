package gitlab_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sk000f/metrix/pkg/gitlab"

	"github.com/sk000f/metrix/pkg/metrix"
	"github.com/stretchr/testify/require"
)

var app metrix.App

func TestMain(m *testing.M) {
	// setup
	app.Init()
	app.AddRoutes(gitlab.InitRoutes())

	os.Exit(m.Run())

	// teardown
}

func TestGitLabHookHeaders(t *testing.T) {

	assert := require.New(t)

	tests := []struct {
		name         string
		headers      http.Header
		responseCode int
	}{
		{
			"GitLab hook received with no event header",
			http.Header{},
			http.StatusBadRequest,
		},
		{
			"GitLab hook received with valid event header for Pipeline event",
			http.Header{
				gitlab.HookHeader: []string{gitlab.PipelineEvent},
			},
			http.StatusOK,
		},
		{
			"GitLab hook received with invalid event header",
			http.Header{
				gitlab.HookHeader: []string{"invalid hook"},
			},
			http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodPost, "/hook/gitlab", nil)
			assert.NoError(err)

			req.Header = tt.headers

			resp := executeRequest(req)

			assert.Equal(tt.responseCode, resp.Code)
		})
	}
}

func TestGitLabHookPipelineEvent(t *testing.T) {
	assert := require.New(t)

	tests := []struct {
		name         string
		headers      http.Header
		dataPath     string
		err          metrix.Error
		responseCode int
	}{
		{
			"GitLab pipeline event with valid data",
			http.Header{
				gitlab.HookHeader: []string{gitlab.PipelineEvent},
			},
			"./testdata/pipelineEvent.json",
			"",
			http.StatusOK,
		},
		{
			"GitLab pipeline event with empty data",
			http.Header{
				gitlab.HookHeader: []string{gitlab.PipelineEvent},
			},
			"./testdata/invalidEvent.json",
			gitlab.ErrMissingGitlabHookData,
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			data, err := os.Open(tt.dataPath)
			assert.NoError(err)
			defer func() {
				_ = data.Close()
			}()

			req, err := http.NewRequest(http.MethodPost, "/hook/gitlab", data)
			assert.NoError(err)

			req.Header = tt.headers

			resp := executeRequest(req)

			assert.Equal(tt.responseCode, resp.Code)
		})
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

package gitlab_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sk000f/metrix/pkg/gitlab"
	"github.com/stretchr/testify/require"
)

var app gitlab.App

func TestMain(m *testing.M) {
	// setup
	app.Init()

	os.Exit(m.Run())

	// teardown
}

func TestGitLabHooks(t *testing.T) {
	t.Run("test initial GitLab endpoint", func(t *testing.T) {

		assert := require.New(t)

		req, err := http.NewRequest(http.MethodPost, "/hook/gitlab", nil)
		assert.NoError(err)

		resp := executeRequest(req)

		assert.NoError(err)
		assert.Equal(http.StatusOK, resp.Code)
	})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

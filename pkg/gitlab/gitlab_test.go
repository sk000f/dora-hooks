package gitlab_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sk000f/metrix/pkg/gitlab"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// setup
	os.Exit(m.Run())
	// teardown
}

func TestGitLabHooks(t *testing.T) {
	t.Run("test initial GitLab endpoint", func(t *testing.T) {

		assert := require.New(t)

		req, err := http.NewRequest(http.MethodPost, "/hook/gitlab", nil)
		assert.NoError(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(gitlab.HookHandler)
		handler.ServeHTTP(rr, req)

		assert.NoError(err)
		assert.Equal(http.StatusOK, rr.Code)
	})
}

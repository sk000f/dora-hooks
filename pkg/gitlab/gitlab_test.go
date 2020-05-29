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

		srv := createTestServer("/hook/gitlab")

		defer srv.Close()

		req, err := http.NewRequest(http.MethodPost, srv.URL+"/hook/gitlab", nil)
		assert.NoError(err)

		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(err)
		assert.Equal(http.StatusOK, resp.StatusCode)
	})
}

func createTestServer(path string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, gitlab.HookHandler)
	return httptest.NewServer(mux)
}

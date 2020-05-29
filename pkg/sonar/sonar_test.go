package sonar_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sk000f/metrix/pkg/sonar"

	"github.com/sk000f/metrix/pkg/metrix"
	"github.com/stretchr/testify/require"
)

var app metrix.App

func TestMain(m *testing.M) {
	// setup
	app.Init()
	app.AddRoutes(sonar.InitRoutes())

	os.Exit(m.Run())

	// teardown
}

func TestSonarHooks(t *testing.T) {
	t.Run("test base sonar hook endpoint", func(t *testing.T) {

		assert := require.New(t)

		req, err := http.NewRequest(http.MethodPost, "/hook/sonar", nil)
		assert.NoError(err)

		resp := executeRequest(req)

		assert.Equal(http.StatusOK, resp.Code)
	})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

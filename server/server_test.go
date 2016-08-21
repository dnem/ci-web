package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	. "github.com/dnem/ci-web/server"
)

var (
	server   *negroni.Negroni
	recorder *httptest.ResponseRecorder
)

// TestHomeHandlerWithoutEnv verifies root handler is functioning correctly
func TestHomeHandlerWithoutEnv(t *testing.T) {
	server = NewServer()

	getHomePageRequest, _ := http.NewRequest("GET", "/", nil)
	recorder = httptest.NewRecorder()
	server.ServeHTTP(recorder, getHomePageRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected response code to be %d, received: %d",
			http.StatusOK, recorder.Code)
	}
	if !strings.Contains(recorder.Body.String(), "stranger") {
		t.Errorf("Expected page to contain `stranger`.")
	}
}

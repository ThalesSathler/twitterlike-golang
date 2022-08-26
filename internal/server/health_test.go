package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/server"
)

func TestGETHealth(t *testing.T) {

	cfg := &server.Config{Port: ":8080"}

	r := server.NewServer(cfg, nil, nil, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	r.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "health", w.Body.String())
}

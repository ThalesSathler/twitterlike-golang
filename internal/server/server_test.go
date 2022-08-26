package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/server"
)

func TestNewServer(t *testing.T) {
	wantedAddr := ":8080"
	cfg := &server.Config{Port: ":8080"}
	got := server.NewServer(cfg, nil, nil, nil)

	assert.Equal(t, wantedAddr, got.Addr)
}

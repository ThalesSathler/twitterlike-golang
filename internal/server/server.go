package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalessathler/twitterlike/internal/auth"
	"github.com/thalessathler/twitterlike/internal/tweet"
	"github.com/thalessathler/twitterlike/internal/user"
)

type commonservice struct {
	router         *gin.Engine
	twitterService *tweet.Service
	userService    *user.Service
	authService    *auth.Service
}

type route interface {
	registry()
}

type Config struct {
	Port string
}

func NewServer(cfg *Config, tw *tweet.Service,
	us *user.Service, auth *auth.Service) *http.Server {

	r := gin.Default()

	service := &commonservice{
		router:         r,
		twitterService: tw,
		userService:    us,
		authService:    auth,
	}

	routes := []route{
		(*HealthRoute)(service),
		(*TweetHandler)(service),
		(*UserHandler)(service),
	}

	for _, route := range routes {
		route.registry()
	}

	return &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}
}

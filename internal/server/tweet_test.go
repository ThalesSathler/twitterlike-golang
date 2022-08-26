package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/auth"
	"github.com/thalessathler/twitterlike/internal/server"
	"github.com/thalessathler/twitterlike/internal/server/fake"
	"github.com/thalessathler/twitterlike/internal/tweet"
)

func TestCreateTweet(t *testing.T) {
	cfg := &server.Config{}

	fakeTweetRepository := fake.New(nil)
	fakeTweetSvc, err := tweet.New(fakeTweetRepository, nil)
	if err != nil {
		assert.Error(t, err)
	}

	monkey.Patch(server.AuthMiddleware, func(authService *auth.Service) gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Set("UserID", "12345")
			c.Next()
		}
	})
	defer monkey.UnpatchAll()

	r := server.NewServer(cfg, fakeTweetSvc, nil, nil)

	payload := `
	{
		"content": "randomContent"
	}
	`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/tweet", bytes.NewReader([]byte(payload)))
	req.Header.Add("Content-Type", "application/json")
	r.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		assert.Error(t, err)
	}

	resp := &tweet.Tweet{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, resp.Content, "randomContent")
	assert.Equal(t, resp.UserID, "12345")
	assert.NotNil(t, resp.ID)
}

func TestCreateTweetErrors(t *testing.T) {
	t.Run("When it returns random database error should return 500", func(t *testing.T) {
		cfg := &server.Config{}

		fakeTweetRepository := fake.New(errors.New("randomDatabaseError"))
		fakeTweetSvc, err := tweet.New(fakeTweetRepository, nil)
		if err != nil {
			assert.Error(t, err)
		}

		monkey.Patch(server.AuthMiddleware, func(authService *auth.Service) gin.HandlerFunc {
			return func(c *gin.Context) {
				c.Set("UserID", "12345")
				c.Next()
			}
		})
		defer monkey.UnpatchAll()

		r := server.NewServer(cfg, fakeTweetSvc, nil, nil)

		payload := `
			{
				"content": "randomContent"
			}
		`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tweet", bytes.NewReader([]byte(payload)))
		req.Header.Add("Content-Type", "application/json")
		r.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, `{"error":"randomDatabaseError"}`, w.Body.String())
	})
	t.Run("When payload is wrong should return 400", func(t *testing.T) {
		cfg := &server.Config{}

		fakeTweetRepository := fake.New(nil)
		fakeTweetSvc, err := tweet.New(fakeTweetRepository, nil)
		if err != nil {
			assert.Error(t, err)
		}

		monkey.Patch(server.AuthMiddleware, func(authService *auth.Service) gin.HandlerFunc {
			return func(c *gin.Context) {
				c.Set("UserID", "12345")
				c.Next()
			}
		})
		defer monkey.UnpatchAll()

		r := server.NewServer(cfg, fakeTweetSvc, nil, nil)

		payload := `
			{
				"randomkey": "randomvalue"
			}
		`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tweet", bytes.NewReader([]byte(payload)))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "12345")
		r.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

package server_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/server"
	"github.com/thalessathler/twitterlike/internal/server/fake"
	"github.com/thalessathler/twitterlike/internal/user"
)

func TestPostUser(t *testing.T) {
	cfg := &server.Config{}

	userRepo := fake.NewFakeUserService(nil)
	userSvc := user.New(userRepo)

	r := server.NewServer(cfg, nil, userSvc, nil)
	t.Run("When everything goes ok", func(t *testing.T) {
		payload := `
			{
				"email": "email@email.com",
				"name": "random name",
				"password": "randompass"
			}
		`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(payload)))
		req.Header.Add("Content-Type", "application/json")
		r.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		body, err := ioutil.ReadAll(w.Body)
		if err != nil {
			assert.Error(t, err)
		}

		resp := &user.User{}
		err = json.Unmarshal(body, resp)
		if err != nil {
			assert.Error(t, err)
		}

		assert.Equal(t, "random name", resp.Name)
		assert.Equal(t, "email@email.com", resp.Email)
		assert.NotEqual(t, "randompass", resp.Password)
		assert.NotEqual(t, "", resp.ID)
	})

	t.Run("When payload is wrong should return 400", func(t *testing.T) {
		payload := `
			{
				"randomkey": "randomvalue"
			}
		`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(payload)))
		req.Header.Add("Content-Type", "application/json")
		r.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

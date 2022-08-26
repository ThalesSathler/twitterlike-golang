package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler commonservice

const (
	userPath = "/user"
	authPath = "/auth"
)

func (t *UserHandler) registry() {
	t.router.POST(userPath, t.PostUser)
	t.router.POST(authPath, t.PostAuth)
}

type PostUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostAuth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (t *UserHandler) PostUser(c *gin.Context) {
	newUser := PostUser{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := t.userService.CreateUser(c.Request.Context(),
		newUser.Name,
		newUser.Email,
		newUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

func (t *UserHandler) PostAuth(c *gin.Context) {
	auth := PostAuth{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := t.authService.Login(c.Request.Context(), auth.Email, auth.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetHandler commonservice

const path = "/tweet"

func (t *TweetHandler) registry() {
	t.router.POST(path, AuthMiddleware(t.authService), t.PostTweet)
}

type Tweet struct {
	Content string `json:"content" binding:"required"`
}

func (t *TweetHandler) PostTweet(c *gin.Context) {
	tweet := Tweet{}
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetString("UserID")
	if userID == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "token is necessary"})
		return
	}
	createdTweet, err := t.twitterService.Tweet(c.Request.Context(), tweet.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdTweet)
}

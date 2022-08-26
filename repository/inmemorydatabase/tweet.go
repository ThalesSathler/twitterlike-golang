package inmemorydatabase

import (
	"context"

	"github.com/google/uuid"
	"github.com/thalessathler/twitterlike/internal/tweet"
)

type TweetRepository struct {
	tweets map[string]*tweet.Tweet
}

func NewTweetRepository() *TweetRepository {
	return &TweetRepository{
		tweets: make(map[string]*tweet.Tweet),
	}
}

func (f *TweetRepository) Create(ctx context.Context, tt *tweet.Tweet) (*tweet.Tweet, error) {
	newTweet := &tweet.Tweet{
		ID:      uuid.NewString(),
		Content: tt.Content,
		UserID:  tt.UserID,
	}

	f.tweets[newTweet.ID] = newTweet
	return newTweet, nil
}

func (f *TweetRepository) Delete(ctx context.Context, tweetID string) error {
	f.tweets[tweetID] = nil
	return nil
}

package tweet

import (
	"context"
	"time"
)

type Tweet struct {
	ID        string
	Content   string
	UserID    string
	CreatedAt time.Time
}

type Retweet struct {
	ID          string
	RetweetedID string
	UserID      string
	Content     string
	CreatedAT   time.Time
}

type TweetRepository interface {
	Create(ctx context.Context, tweet *Tweet) (*Tweet, error)
	Delete(ctx context.Context, tweetID string) error
}

type RetweetRepository interface {
	Create(ctx context.Context, retweet *Retweet) (*Retweet, error)
	Delete(ctx context.Context, tweetID string) error
}

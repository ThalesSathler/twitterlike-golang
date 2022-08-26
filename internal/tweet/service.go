package tweet

import (
	"time"

	"context"
)

type Service struct {
	tweetRepo   TweetRepository
	retweetRepo RetweetRepository
}

func New(tweetRepo TweetRepository, retweetRepo RetweetRepository) (*Service, error) {
	return &Service{
		tweetRepo:   tweetRepo,
		retweetRepo: retweetRepo,
	}, nil
}

func (s *Service) Tweet(ctx context.Context, content, userID string) (*Tweet, error) {
	if content == "" {
		return nil, ErrMissingContent
	}
	if userID == "" {
		return nil, ErrEmptyUserID
	}

	tweet := &Tweet{
		Content:   content,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	return s.tweetRepo.Create(ctx, tweet)
}

func (s *Service) Retweet(ctx context.Context, tweetID, content, userID string) (*Retweet, error) {
	if content == "" {
		return nil, ErrMissingContent
	}
	if userID == "" {
		return nil, ErrEmptyUserID
	}
	if tweetID == "" {
		return nil, ErrEmptyTweetID
	}

	retweet := &Retweet{
		RetweetedID: tweetID,
		UserID:      userID,
		Content:     content,
		CreatedAT:   time.Now(),
	}

	return s.retweetRepo.Create(ctx, retweet)
}

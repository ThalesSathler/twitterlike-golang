package fake

import (
	"context"

	"github.com/google/uuid"
	"github.com/thalessathler/twitterlike/internal/tweet"
)

type FakeTwitterService struct {
	tweets        map[string]*tweet.Tweet
	expectedError error
}

func New(err error) *FakeTwitterService {
	return &FakeTwitterService{
		expectedError: err,
		tweets:        make(map[string]*tweet.Tweet),
	}
}

func (f *FakeTwitterService) Create(ctx context.Context, tt *tweet.Tweet) (*tweet.Tweet, error) {
	if f.expectedError != nil {
		return nil, f.expectedError
	}

	newTweet := &tweet.Tweet{
		ID:      uuid.NewString(),
		Content: tt.Content,
		UserID:  tt.UserID,
	}

	f.tweets[newTweet.ID] = newTweet
	return newTweet, nil
}

func (f *FakeTwitterService) Delete(ctx context.Context, tweetID string) error {
	if f.expectedError != nil {
		return f.expectedError
	}

	f.tweets[tweetID] = nil
	return nil
}

package tweet_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/thalessathler/twitterlike/internal/tweet"
	"github.com/thalessathler/twitterlike/internal/tweet/fake"
)

type givenTweet struct {
	content string
	userID  string
	tweetID string
}

func TestTweet(t *testing.T) {
	tests := []struct {
		name      string
		given     *givenTweet
		want      *tweet.Tweet
		wantedErr error
	}{
		{
			name: "When tweet has no body",
			given: &givenTweet{
				content: "",
				userID:  "somerandomID",
			},
			wantedErr: tweet.ErrMissingContent,
		},
		{
			name: "When tweet has no userID",
			given: &givenTweet{
				content: "some random content",
				userID:  "",
			},
			wantedErr: tweet.ErrEmptyUserID,
		},
	}

	ctx := context.Background()

	tweetRepo := fake.NewTweetRepository(t)
	retweetRepo := fake.NewRetweetRepository(t)
	service, err := tweet.New(tweetRepo, retweetRepo)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Tweet(ctx, tt.given.content, tt.given.userID)
			assert.Nil(t, got)
			assert.Equal(t, tt.wantedErr, err)
		})
	}

	t.Run("When random error occurs", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid"}
		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedTweet := &tweet.Tweet{
			Content:   "somerandomcontent",
			UserID:    "somerandomid",
			CreatedAt: mockedTime,
		}

		tweetRepo.
			On("Create", ctx, wantedTweet).
			Return(nil, errors.New("SomeRandomError")).
			Once()

		got, err := service.Tweet(ctx, given.content, given.userID)
		assert.Nil(t, got)
		assert.Equal(t, "SomeRandomError", err.Error())
	})

	t.Run("When context times out", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid"}
		ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
		cancel()
		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedTweet := &tweet.Tweet{
			Content:   "somerandomcontent",
			UserID:    "somerandomid",
			CreatedAt: mockedTime,
		}

		tweetRepo.
			On("Create", ctx, wantedTweet).
			Return(nil, context.DeadlineExceeded).
			Once()

		got, err := service.Tweet(ctx, given.content, given.userID)
		assert.Nil(t, got)
		assert.Equal(t, context.DeadlineExceeded, err)
	})

	t.Run("When everything goes ok", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid"}
		response := &tweet.Tweet{ID: "SomerandomID", Content: "somerandomcontent", UserID: "somerandomid"}
		wantedResponse := &tweet.Tweet{ID: "SomerandomID", Content: "somerandomcontent", UserID: "somerandomid"}

		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedTweet := &tweet.Tweet{
			Content:   "somerandomcontent",
			UserID:    "somerandomid",
			CreatedAt: mockedTime,
		}

		tweetRepo.
			On("Create", ctx, wantedTweet).
			Return(response, nil).
			Once()

		got, err := service.Tweet(ctx, given.content, given.userID)
		assert.Equal(t, wantedResponse, got)
		assert.Nil(t, err)
	})
}

func TestRetweet(t *testing.T) {
	tests := []struct {
		name      string
		given     *givenTweet
		wantedErr error
	}{
		{
			name: "When tweet has no body",
			given: &givenTweet{
				content: "",
				userID:  "somerandomID",
				tweetID: "somerandomID",
			},
			wantedErr: tweet.ErrMissingContent,
		},
		{
			name: "When tweet has no userID",
			given: &givenTweet{
				content: "some random content",
				userID:  "",
				tweetID: "somerandomID",
			},
			wantedErr: tweet.ErrEmptyUserID,
		},
		{
			name: "When tweet has no tweetID",
			given: &givenTweet{
				content: "some random content",
				userID:  "somerandomID",
				tweetID: "",
			},
			wantedErr: tweet.ErrEmptyTweetID,
		},
	}

	ctx := context.Background()

	tweetRepo := fake.NewTweetRepository(t)
	retweetRepo := fake.NewRetweetRepository(t)
	service, err := tweet.New(tweetRepo, retweetRepo)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Retweet(ctx, tt.given.tweetID, tt.given.content, tt.given.userID)
			assert.Nil(t, got)
			assert.Equal(t, tt.wantedErr, err)
		})
	}

	t.Run("When random error occurs", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid", tweetID: "somerandomid"}

		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedRetweet := &tweet.Retweet{
			RetweetedID: "somerandomid",
			UserID:      "somerandomid",
			Content:     "somerandomcontent",
			CreatedAT:   mockedTime,
		}

		retweetRepo.
			On("Create", ctx, wantedRetweet).
			Return(nil, errors.New("SomeRandomError")).
			Once()

		got, err := service.Retweet(ctx, given.tweetID, given.content, given.userID)
		assert.Nil(t, got)
		assert.Equal(t, "SomeRandomError", err.Error())
	})

	t.Run("When context times out", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid", tweetID: "somerandomid"}
		ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
		cancel()

		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedRetweet := &tweet.Retweet{
			RetweetedID: "somerandomid",
			UserID:      "somerandomid",
			Content:     "somerandomcontent",
			CreatedAT:   mockedTime,
		}

		retweetRepo.
			On("Create", ctx, wantedRetweet).
			Return(nil, context.DeadlineExceeded).
			Once()

		got, err := service.Retweet(ctx, given.tweetID, given.content, given.userID)
		assert.Nil(t, got)
		assert.Equal(t, context.DeadlineExceeded, err)
	})

	t.Run("When everything goes ok", func(t *testing.T) {
		given := givenTweet{content: "somerandomcontent", userID: "somerandomid", tweetID: "somerandomid"}
		response := &tweet.Retweet{ID: "SomerandomID", Content: "somerandomcontent", UserID: "somerandomid", RetweetedID: "somerandomid"}
		wantedResponse := &tweet.Retweet{ID: "SomerandomID", Content: "somerandomcontent", UserID: "somerandomid", RetweetedID: "somerandomid"}

		mockedTime := time.Now()

		monkey.Patch(time.Now, func() time.Time {
			return mockedTime
		})
		defer monkey.UnpatchAll()

		wantedRetweet := &tweet.Retweet{
			RetweetedID: "somerandomid",
			UserID:      "somerandomid",
			Content:     "somerandomcontent",
			CreatedAT:   mockedTime,
		}

		retweetRepo.
			On("Create", ctx, wantedRetweet).
			Return(response, nil).
			Once()

		got, err := service.Retweet(ctx, given.tweetID, given.content, given.userID)
		assert.Equal(t, wantedResponse, got)
		assert.Nil(t, err)
	})
}

type givenRepo struct {
	tweetRepo   tweet.TweetRepository
	retweetRepo tweet.RetweetRepository
}

func TestNew(t *testing.T) {
	tweetRepo := fake.NewTweetRepository(t)
	retweetRepo := fake.NewRetweetRepository(t)
	tests := []struct {
		name  string
		given givenRepo
	}{
		{
			name:  "when tweetRepo is nil",
			given: givenRepo{nil, retweetRepo},
		},
		{
			name:  "when retweetRepo is nil",
			given: givenRepo{tweetRepo, nil},
		},
		{
			name:  "when everything ok",
			given: givenRepo{tweetRepo, retweetRepo},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tweet.New(tt.given.tweetRepo, tt.given.retweetRepo)
			assert.NoError(t, err)
		})
	}
}

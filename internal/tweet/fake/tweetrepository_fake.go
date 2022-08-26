// Code generated by mockery v2.12.3. DO NOT EDIT.

package fake

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	"github.com/thalessathler/twitterlike/internal/tweet"
)

// TweetRepository is an autogenerated mock type for the TweetRepository type
type TweetRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *TweetRepository) Create(ctx context.Context, _a1 *tweet.Tweet) (*tweet.Tweet, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *tweet.Tweet
	if rf, ok := ret.Get(0).(func(context.Context, *tweet.Tweet) *tweet.Tweet); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tweet.Tweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *tweet.Tweet) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, tweetID
func (_m *TweetRepository) Delete(ctx context.Context, tweetID string) error {
	ret := _m.Called(ctx, tweetID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tweetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewTweetRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTweetRepository creates a new instance of TweetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTweetRepository(t NewTweetRepositoryT) *TweetRepository {
	mock := &TweetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.12.3. DO NOT EDIT.

package fake

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	"github.com/thalessathler/twitterlike/internal/tweet"
)

// RetweetRepository is an autogenerated mock type for the RetweetRepository type
type RetweetRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, retweet
func (_m *RetweetRepository) Create(ctx context.Context, retweet *tweet.Retweet) (*tweet.Retweet, error) {
	ret := _m.Called(ctx, retweet)

	var r0 *tweet.Retweet
	if rf, ok := ret.Get(0).(func(context.Context, *tweet.Retweet) *tweet.Retweet); ok {
		r0 = rf(ctx, retweet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tweet.Retweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *tweet.Retweet) error); ok {
		r1 = rf(ctx, retweet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, tweetID
func (_m *RetweetRepository) Delete(ctx context.Context, tweetID string) error {
	ret := _m.Called(ctx, tweetID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tweetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewRetweetRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRetweetRepository creates a new instance of RetweetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRetweetRepository(t NewRetweetRepositoryT) *RetweetRepository {
	mock := &RetweetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
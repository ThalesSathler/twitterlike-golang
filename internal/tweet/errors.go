package tweet

import "errors"

var (
	ErrMissingContent = errors.New("content is missing")
	ErrEmptyUserID    = errors.New("userID is missing")
	ErrEmptyTweetID   = errors.New("retweetID is missing")
)

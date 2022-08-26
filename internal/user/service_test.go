package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/user"
	"github.com/thalessathler/twitterlike/internal/user/fake"
)

func Test_AuthorizeUser(t *testing.T) {
	userRepo := fake.NewUserRepository(t)
	svc := user.New(userRepo)
	t.Run("When user is not found", func(t *testing.T) {
		givenEmail, givenPassword := "randomemail", "hashedPassword"

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(nil, user.ErrNotFound).
			Once()

		got, err := svc.AuthorizeUser(context.Background(), givenEmail, givenPassword)

		assert.Nil(t, got)
		assert.ErrorIs(t, err, user.ErrNotFound)
	})
	t.Run("When password is wrong", func(t *testing.T) {
		givenEmail, givenPassword := "randomemail", "hashedPassword"

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(&user.User{Password: "wrongPassword"}, nil).
			Once()

		got, err := svc.AuthorizeUser(context.Background(), givenEmail, givenPassword)

		assert.Nil(t, got)
		assert.ErrorIs(t, err, user.ErrWrongCredentials)
	})
	t.Run("When db times out", func(t *testing.T) {
		givenEmail, givenPassword := "randomemail", "hashedPassword"

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(nil, context.DeadlineExceeded).
			Once()

		got, err := svc.AuthorizeUser(context.Background(), givenEmail, givenPassword)

		assert.Nil(t, got)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
	t.Run("When everything goes ok", func(t *testing.T) {
		givenEmail, givenPassword := "random@email.com", "hashedPassword"

		wantedUser := &user.User{
			ID:       "12345",
			Name:     "Some Random Name",
			Email:    "random@email.com",
			Password: "hashedPassword",
		}

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(wantedUser, nil).
			Once()

		got, err := svc.AuthorizeUser(context.Background(), givenEmail, givenPassword)

		assert.Equal(t, wantedUser, got)
		assert.Nil(t, err)
	})
}

func Test_CreateUser(t *testing.T) {
	userRepo := fake.NewUserRepository(t)
	svc := user.New(userRepo)
	t.Run("When db returns error", func(t *testing.T) {

		givenName, givenEmail, givenPassword := "name", "email", "password"

		givenUser := &user.User{
			Name:     givenName,
			Email:    givenEmail,
			Password: "hashedPassword",
		}

		userRepo.
			On("Create", context.Background(), givenUser).
			Return(nil, errors.New("some db error")).
			Once()

		ctx := context.Background()

		got, err := svc.CreateUser(ctx, givenName, givenEmail, givenPassword)

		assert.Nil(t, got)
		assert.Equal(t, "some db error", err.Error())
	})
	t.Run("When user is created", func(t *testing.T) {

		givenName, givenEmail, givenPassword := "name", "email", "password"

		givenUser := &user.User{
			Name:     givenName,
			Email:    givenEmail,
			Password: "hashedPassword",
		}

		createdUser := &user.User{
			ID:       "12345",
			Name:     "name",
			Email:    "email",
			Password: "",
		}

		userRepo.
			On("Create", context.Background(), givenUser).
			Return(createdUser, nil).
			Once()

		ctx := context.Background()

		got, err := svc.CreateUser(ctx, givenName, givenEmail, givenPassword)

		assert.Equal(t, createdUser, got)
		assert.Nil(t, err)
	})
}

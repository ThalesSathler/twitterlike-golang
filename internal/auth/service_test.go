package auth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thalessathler/twitterlike/internal/auth"
	"github.com/thalessathler/twitterlike/internal/user"
	"github.com/thalessathler/twitterlike/internal/user/fake"
)

func Test_Login(t *testing.T) {
	userRepo := fake.NewUserRepository(t)
	usersvc := user.New(userRepo)

	authsvc := auth.New(usersvc)
	t.Run("When everything goes ok", func(t *testing.T) {
		givenEmail, givenPassword := "random@email.com", "hashedPassword"

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(&user.User{
				ID:       "someID",
				Name:     "randomName",
				Email:    "random@email.com",
				Password: "hashedPassword",
			}, nil).
			Once()

		token, err := authsvc.Login(context.Background(),
			givenEmail,
			givenPassword,
		)

		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})
	t.Run("When something goes wrong", func(t *testing.T) {
		givenEmail, givenPassword := "random@email.com", "hashedPassword"

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(&user.User{Password: "wrong"}, nil).
			Once()

		token, err := authsvc.Login(context.Background(),
			givenEmail,
			givenPassword,
		)

		assert.Empty(t, token)
		assert.ErrorIs(t, err, user.ErrWrongCredentials)
	})
}

func Test_ValidateToken(t *testing.T) {
	userRepo := fake.NewUserRepository(t)
	usersvc := user.New(userRepo)

	authsvc := auth.New(usersvc)
	t.Run("When everything goes ok", func(t *testing.T) {
		givenEmail, givenPassword := "random@email.com", "hashedPassword"

		wantedUser := &user.User{
			ID:    "someID",
			Name:  "randomName",
			Email: "random@email.com",
		}

		userRepo.
			On("GetByEmail", context.Background(), givenEmail).
			Return(&user.User{
				ID:       "someID",
				Name:     "randomName",
				Email:    "random@email.com",
				Password: "hashedPassword",
			}, nil).
			Once()

		token, err := authsvc.Login(context.Background(),
			givenEmail,
			givenPassword,
		)

		validatedUser, err := authsvc.ValidateToken(context.Background(), token)

		assert.NotEmpty(t, token)
		assert.Nil(t, err)
		assert.Equal(t, wantedUser, validatedUser)
	})
	t.Run("When token is not valid", func(t *testing.T) {
		token := "eyJhbGciOJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InNvbWVJRCIsIm5hbWUiOiJyYW5kb21OYW1lIiwiZW1haWwiOiJyYW5kb21AZW1haWwuY29tIiwiZXhwIjoxNjYxNTM2ODQ3fQ.WuGMRWQ8AeFo9128NfA5ADpEzkqO9skMsqbNFwG3yds"

		validatedUser, err := authsvc.ValidateToken(context.Background(), token)

		assert.Error(t, err)
		assert.Nil(t, validatedUser)
	})
}

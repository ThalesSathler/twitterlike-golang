package inmemorydatabase

import (
	"context"

	"github.com/google/uuid"
	"github.com/thalessathler/twitterlike/internal/user"
)

type UserRepository struct {
	users         map[string]*user.User
	expectedError error
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*user.User),
	}
}

func (f *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	if f.expectedError != nil {
		return nil, f.expectedError
	}

	newUser := &user.User{
		ID:       uuid.NewString(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	f.users[newUser.Email] = newUser
	return newUser, nil
}

func (f *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	if f.expectedError != nil {
		return nil, f.expectedError
	}

	if user, ok := f.users[email]; ok {
		return user, nil
	}

	return nil, user.ErrNotFound
}

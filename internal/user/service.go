package user

import "context"

type Service struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, name, email, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: hashPass(password),
	}
	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	clonedUser := *createdUser
	clonedUser.Password = ""
	return &clonedUser, nil
}

func (s *Service) AuthorizeUser(ctx context.Context, email, password string) (*User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !isValidPassword(password, user.Password) {
		return nil, ErrWrongCredentials
	}
	user.Password = ""
	return user, nil
}

func isValidPassword(givenPassword, storedPassword string) bool {
	return hashPass(givenPassword) == storedPassword
}

//TODO Create HashedPassword
func hashPass(p string) string {
	return "hashedPassword"
}

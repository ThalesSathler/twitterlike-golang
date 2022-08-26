package auth

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thalessathler/twitterlike/internal/user"
)

type Service struct {
	userService *user.Service
}

type claims struct {
	user.User
	jwt.StandardClaims
}

//TODO Get Value from env var
var jwtKey = []byte("gangnamstyle")

func New(userSvc *user.Service) *Service {
	return &Service{
		userService: userSvc,
	}
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userService.AuthorizeUser(ctx, email, password)
	if err != nil {
		return "", err
	}

	return tokenize(user)
}

func tokenize(user *user.User) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := claims{
		*user,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *Service) ValidateToken(ctx context.Context, token string) (*user.User, error) {
	claims := &claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, ErrInvalidToken
	}

	return &claims.User, nil
}

package usecase

import (
	"strconv"
	"time"

	"github.com/FernandoCagale/serverless-auth/pkg/auth"
	"github.com/FernandoCagale/serverless-auth/pkg/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

//Service interface
type Service struct {
	repo auth.Repository
}

//NewService create new service
func NewService(r auth.Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Login auth
func (s *Service) Login(e *entity.Auth) (string, error) {
	err := e.Validate()
	if err != nil {
		return "", entity.ErrInvalidPayload
	}

	err = s.repo.Login(e)
	if err != nil {
		return "", err
	}

	token, err := createJwtToken(e)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createJwtToken(e *entity.Auth) (string, error) {
	claims := auth.JwtClaims{
		ID:    e.ID,
		Scope: []string{"admin"},
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(e.ID),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return rawToken.SignedString([]byte("secret"))
}

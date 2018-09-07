package auth

import "github.com/FernandoCagale/serverless-auth/pkg/entity"

//UseCase use case interface
type UseCase interface {
	Login(auth *entity.Auth) (string, error)
}

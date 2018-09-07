package auth

import "github.com/FernandoCagale/serverless-auth/pkg/entity"

//Repository repository interface
type Repository interface {
	Login(auth *entity.Auth) (err error)
}

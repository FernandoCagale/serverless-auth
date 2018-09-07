package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

//Auth entity
type Auth struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//ValidatePassword validate password
func (e *Auth) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
}

//Validate struct
func (e Auth) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Username, validation.Required, validation.Length(5, 20)),
		validation.Field(&e.Password, validation.Required, validation.Length(3, 10)),
	)
}

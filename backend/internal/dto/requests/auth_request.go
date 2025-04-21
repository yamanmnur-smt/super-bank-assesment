package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r Login) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type Register struct {
	Name     string
	Username string
	Password string
}

func (r Register) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

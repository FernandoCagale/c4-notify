package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Notify struct {
	Code   string   `json:"code" bson:"_id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Phone  string   `json:"phone"`
	Notify []string `json:"notify"`
}

func (e Notify) Validate() error {
	return validation.ValidateStruct(&e,
	)
}

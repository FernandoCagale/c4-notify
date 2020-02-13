package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Customer struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Phone  string   `json:"phone"`
	Notify []string `json:"notify"`
}

func (e Customer) Validate() error {
	return validation.ValidateStruct(&e,
	)
}

func (e Customer) ToNotify() *Notify {
	return &Notify{
		Code:   e.Code,
		Name:   e.Name,
		Email:  e.Email,
		Phone:  e.Phone,
		Notify: e.Notify,
	}
}
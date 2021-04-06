package model

import (
	"errors"
	"time"
)

type User struct {
	ID          string    `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	BirthDate   time.Time `json:"birth_date"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

func (u *User) ValidCreatePayload() error {
	if len(u.FirstName) == 0 {
		return errors.New("empty first name")
	}
	return nil
}

func (u *User) ValidUpdatePayload() error {
	return nil
}

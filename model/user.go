package model

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	DateFromat = "2006-01-02"
)

type User struct {
	ID          string    `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	BirthDate   time.Time `json:"birth_date"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

type UserAux struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	var aux UserAux

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	u.FirstName = aux.FirstName
	u.LastName = aux.LastName
	u.Email = aux.Email
	u.PhoneNumber = aux.PhoneNumber
	u.BirthDate, err = time.Parse(DateFromat, aux.BirthDate)
	if err != nil {
		return err
	}

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	var aux UserAux
	aux.FirstName = u.FirstName
	aux.LastName = u.LastName
	aux.Email = u.Email
	aux.PhoneNumber = u.PhoneNumber
	aux.BirthDate = u.BirthDate.Format(DateFromat)
	return json.Marshal(aux)
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

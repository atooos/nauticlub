package model

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DateFromat = "2006-01-02"
)

type MembershipType int

const (
	MembershipUnknown MembershipType = iota
	MembershipFamilly
	MembershipIndividual
)

var tblMembership = []string{
	"unknown",
	"familly",
	"individual",
}

func (mt MembershipType) String() string {
	return tblMembership[mt]
}

type User struct {
	ID          string         `json:"uuid"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	BirthDate   time.Time      `json:"birth_date"`
	Email       string         `json:"email"`
	Password    string         `json:"pass"`
	PhoneNumber string         `json:"phone_number"`
	Membership  MembershipType `json:"membership"`
}

type UserAux struct {
	ID          string `json:"uuid"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	Password    string `json:"pass,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Membership  string `json:"membership,omitempty"`
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
	if len(aux.Password) > 0 {
		h := sha256.New()
		h.Write([]byte(aux.Password))
		u.Password = fmt.Sprintf("%x", h.Sum(nil))
	}
	if len(aux.BirthDate) != 0 {
		u.BirthDate, err = time.Parse(DateFromat, aux.BirthDate)
		if err != nil {
			return err
		}
	}

	if len(aux.Membership) != 0 {
		switch aux.Membership {
		case "familly":
			u.Membership = MembershipFamilly
		case "individual":
			u.Membership = MembershipIndividual
		default:
			u.Membership = MembershipUnknown
		}
	}

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	var aux UserAux
	aux.ID = u.ID
	aux.FirstName = u.FirstName
	aux.LastName = u.LastName
	aux.Email = u.Email
	aux.PhoneNumber = u.PhoneNumber
	aux.BirthDate = u.BirthDate.Format(DateFromat)
	aux.Password = ""
	aux.Membership = u.Membership.String()
	return json.Marshal(aux)
}

func (u *User) ValidCreatePayload() error {
	if len(u.FirstName) == 0 {
		return errors.New("empty first name")
	}
	if len(u.LastName) == 0 {
		return errors.New("empty first name")
	}
	if len(u.Email) == 0 {
		return errors.New("empty first name")
	}
	if len(u.FirstName) == 0 {
		return errors.New("empty first name")
	}
	return nil
}

func (u *User) ValidUpdatePayload() error {
	return nil
}

package model

import (
	"encoding/json"
	"time"
)

type Sub struct {
	ID        string    `json:"uuid"`
	UserID    string    `json:"user_uuid"`
	Kind      SubType   `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type SubAux struct {
	ID        string `json:"uuid"`
	UserID    string `json:"user_uuid"`
	Kind      string `json:"kind"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at,omitempty"`
	ExpiredAt string `json:"expired_at,omitempty"`
}

func (s *Sub) UnmarshalJSON(b []byte) error {
	var aux SubAux
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	if len(aux.ID) != 0 {
		s.ID = aux.ID
	}

	s.UserID = aux.UserID
	if len(aux.CreatedAt) != 0 {
		s.CreatedAt, err = time.Parse(DateFromat, aux.CreatedAt)
		if err != nil {
			return err
		}
	}

	if len(aux.UpdateAt) != 0 {
		s.UpdateAt, err = time.Parse(DateFromat, aux.UpdateAt)
		if err != nil {
			return err
		}
	}

	if len(aux.ExpiredAt) != 0 {
		s.ExpiredAt, err = time.Parse(DateFromat, aux.ExpiredAt)
		if err != nil {
			return err
		}
	}

	if len(aux.Kind) != 0 {
		switch aux.Kind {
		case "licence":
			s.Kind = SubLicence
		case "club":
			s.Kind = SubClub
		default:
			s.Kind = SubUnknown
		}
	}
	return nil
}

func (s Sub) MarshalJSON() ([]byte, error) {
	var aux SubAux
	aux.ID = s.ID
	aux.CreatedAt = s.CreatedAt.Format(DateFromat)
	aux.ExpiredAt = s.ExpiredAt.Format(DateFromat)
	aux.UpdateAt = s.UpdateAt.Format(DateFromat)
	aux.UserID = s.UserID
	aux.Kind = s.Kind.String()
	return json.Marshal(&aux)
}

type SubType int

const (
	SubUnknown SubType = iota
	SubLicence
	SubClub
)

var tblSub = []string{
	"unknown",
	"licence",
	"club",
}

func (st SubType) String() string {
	return tblSub[st]
}

func (s *Sub) ValidCreatePayload() error {
	return nil
}

func (s *Sub) ValidUpdatePayload() error {
	return nil
}

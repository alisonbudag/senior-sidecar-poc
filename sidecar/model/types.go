package model

import "time"

type AuthType int

const (
	G5 = iota
	G7
	SAML
	LDAP
)

type Role struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type User struct {
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Email       string    `json:"Email"`
	Locale      string    `json:"Locale"`
	AuthType    AuthType  `json:"AuthType"`
	Birthdate   time.Time `json:"Birthdate"`
	Photo       []byte    `json:"Photo"`
	Blocked     bool      `json:"Blocked"`
	Roles       []Role    `json:"Roles"`
}

package storage

import (
	"github.com/google/uuid"
)

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name"`
}

type NewSecret struct {
	Secret string `json:"secret"`
}

type Secret struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Secret    string    `json:"secret"`
	Guessed   bool      `json:"guessed"`
	GuesserId uuid.UUID `json:"guesser"`
}

func (nu NewUser) ToUser() User {
	return User{
		Id:   uuid.New(),
		Name: nu.Name,
	}
}

func (ns NewSecret) ToSecret() Secret {
	return Secret{
		Id:        uuid.New(),
		Secret:    ns.Secret,
		Guessed:   false,
		GuesserId: uuid.UUID{},
	}
}

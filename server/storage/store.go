package storage

import "github.com/google/uuid"

type Store interface {
	HealthCheck() bool

	CreateUser(newUser NewUser) (User, error)
	GetUser(id uuid.UUID) (User, error)
	GetAllUsers() ([]User, error)

	CreateSecret(newSecret NewSecret) (Secret, error)
	GetSecret(id uuid.UUID) (Secret, error)
	GetAllSecrets() ([]Secret, error)
	UpdateSecret(id uuid.UUID, guesser uuid.UUID) (Secret, error)
}

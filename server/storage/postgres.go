package storage

import (
	"secret_api/config"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	Db *gorm.DB
}

func NewPostgresStore(cfg config.Config) (PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl), &gorm.Config{})
	if err != nil {
		return PostgresStore{}, PostgresConnError{Err: err.Error()}
	}

	db.AutoMigrate(&User{}, &Secret{})

	return PostgresStore{Db: db}, nil
}

func (ps *PostgresStore) HealthCheck() bool {
	result := -1
	ps.Db.Raw("SELECT 1;").Scan(&result)

	return result != -1
}

func (ps *PostgresStore) CreateUser(newUser NewUser) (User, error) {
	user := newUser.ToUser()

	result := ps.Db.Create(&user)
	if result.Error != nil {
		return User{}, CreationFailedError{Resource: "user", Err: result.Error.Error()}
	}

	return user, nil
}

func (ps *PostgresStore) GetUser(id uuid.UUID) (User, error) {
	var user User
	result := ps.Db.Find(&user, id)
	if result.Error != nil {
		return User{}, NotFound{Resource: "user", Id: id}
	}
	return user, nil
}

func (ps *PostgresStore) GetAllUsers() ([]User, error) {
	var allUsers []User
	result := ps.Db.Find(&allUsers)
	if result.Error != nil {
		return []User{}, GetAllError{Resource: "users"}
	}

	return allUsers, nil
}

func (ps *PostgresStore) CreateSecret(newSecret NewSecret) (Secret, error) {
	secret := newSecret.ToSecret()

	result := ps.Db.Create(&secret)
	if result.Error != nil {
		return Secret{}, CreationFailedError{Resource: "secret", Err: result.Error.Error()}
	}

	return secret, nil
}

func (ps *PostgresStore) GetSecret(id uuid.UUID) (Secret, error) {
	var secret Secret
	result := ps.Db.Find(&secret, id)
	if result.Error != nil {
		return Secret{}, NotFound{Resource: "secret", Id: id}
	}
	return secret, nil
}

func (ps *PostgresStore) GetAllSecrets() ([]Secret, error) {
	var allSecrets []Secret
	result := ps.Db.Find(&allSecrets)
	if result.Error != nil {
		return []Secret{}, GetAllError{Resource: "secrets"}
	}

	return allSecrets, nil
}

func (ps *PostgresStore) UpdateSecret(id uuid.UUID, guesser uuid.UUID) (Secret, error) {
	secret, err := ps.GetSecret(id)
	if err != nil {
		return Secret{}, nil
	}

	secret.Guessed = true
	secret.GuesserId = guesser

	result := ps.Db.Save(&secret)
	if result.Error != nil {
		return Secret{}, UpdateError{Id: id, Resource: "secret"}
	}

	return secret, nil
}

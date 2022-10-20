package controllers

import (
	"secret_api/storage"
	"testing"

	"github.com/google/uuid"
)

func createSecret(st storage.MemoryStore, t *testing.T) storage.Secret {
	secret, err := CreateSecret(&st, "shhh")
	if err != nil {
		t.Errorf("Failed to create secret: %s", err.Error())
	}

	return secret
}

func compareSecrets(secret1 storage.Secret, secret2 storage.Secret, t *testing.T) {
	if secret1.Secret != secret2.Secret {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.Secret, secret2.Secret)
	}
	if secret1.Id != secret2.Id {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.Id, secret2.Id)
	}
	if secret1.GuesserId.String() != secret2.GuesserId.String() {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.GuesserId, secret2.GuesserId)
	}
	if secret1.Guessed != secret2.Guessed {
		t.Errorf("Secrets values doesn't match. Left: '%t', Right: '%t'", secret1.Guessed, secret2.Guessed)
	}
}

func TestCreateSecret(t *testing.T) {
	st := storage.NewMemoryStore()

	secret_str := "shhh"

	secret := createSecret(st, t)

	if secret.Secret != secret_str {
		t.Errorf("Created secret's secret doesn't match. Want: '%s', Got: '%s'", secret_str, secret.Secret)
	}
}

func TestGetSecret(t *testing.T) {
	st := storage.NewMemoryStore()

	createdSecret := createSecret(st, t)

	gotSecret, err := GetSecret(&st, createdSecret.Id)
	if err != nil {
		t.Errorf("Failed to get secret: %s", err.Error())
	}

	compareSecrets(createdSecret, gotSecret, t)
}

func TestGetAllSecrets(t *testing.T) {
	st := storage.NewMemoryStore()

	createSecret(st, t)
	createSecret(st, t)

	secrets, err := GetAllSecrets(&st)
	if err != nil {
		t.Errorf("Failed to get all secrets: %s", err.Error())
	}

	if len(secrets) != 2 {
		t.Errorf("Len of secrets doesn't match expected value. Want '2', Got '%d'", len(secrets))
	}
}

func TestGuessSecret(t *testing.T) {
	st := storage.NewMemoryStore()

	createdSecret := createSecret(st, t)
	user := createUser(st, t)

	secret, correct, err := GuessSecret(&st, createdSecret.Id, "shhh", user.Id)
	if err != nil {
		t.Errorf("Failed to guess secret: %s", err.Error())
	}

	if !secret.Guessed {
		t.Error("Secret is not correctly 'guessed' after guesing correctly")
	}

	if !correct {
		t.Error("'correct' variable doesn't match expected. Want 'true', Got 'false'")
	}
}

func TestGuessSecretIncorrectly(t *testing.T) {
	st := storage.NewMemoryStore()

	createdSecret := createSecret(st, t)
	user := createUser(st, t)

	_, correct, _ := GuessSecret(&st, createdSecret.Id, "wrong", user.Id)

	if correct {
		t.Error("'correct' variable doens't match expected. Want 'false', Got 'true'")
	}
}

func TestGuessInvalidSecret(t *testing.T) {
	st := storage.NewMemoryStore()

	user := createUser(st, t)

	_, _, err := GuessSecret(&st, uuid.New(), "something", user.Id)
	if err == nil {
		t.Error("Err should not be nil when trying to guess invalid secret")
	}
}

func TestGuessSecretWithInvalidUser(t *testing.T) {
	st := storage.NewMemoryStore()

	secret := createSecret(st, t)

	_, _, err := GuessSecret(&st, secret.Id, "something", uuid.New())
	if err == nil {
		t.Error("Err should not be nil when trying to guess secret with invalid user")
	}
}

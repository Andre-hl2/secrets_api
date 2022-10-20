package controllers

import (
	"secret_api/storage"

	"github.com/google/uuid"
)

func CreateSecret(st storage.Store, secret string) (storage.Secret, error) {
	return st.CreateSecret(storage.NewSecret{Secret: secret})
}

func GetSecret(st storage.Store, id uuid.UUID) (storage.Secret, error) {
	return st.GetSecret(id)
}

func GetAllSecrets(st storage.Store) ([]storage.Secret, error) {
	return st.GetAllSecrets()
}

func GuessSecret(st storage.Store, secret_id uuid.UUID, guess string, guesser_id uuid.UUID) (storage.Secret, bool, error) {
	secret, err := st.GetSecret(secret_id)
	if err != nil {
		return secret, false, err
	}

	if _, err := st.GetUser(guesser_id); err != nil {
		return storage.Secret{}, false, err
	}

	correct := secret.Secret == guess

	if correct {
		secret, err := st.UpdateSecret(secret_id, guesser_id)
		if err != nil {
			return secret, true, err
		}

		return secret, true, nil
	}

	return secret, false, nil

}

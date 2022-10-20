package storage

import "github.com/google/uuid"

type MemoryStore struct {
	users   map[string]User
	secrets map[string]Secret
}

func NewMemoryStore() MemoryStore {
	return MemoryStore{
		users:   map[string]User{},
		secrets: map[string]Secret{},
	}
}

func (ms *MemoryStore) HealthCheck() bool {
	return true
}

func (ms *MemoryStore) CreateUser(newUser NewUser) (User, error) {
	user := newUser.ToUser()
	ms.users[user.Id.String()] = user

	return user, nil
}

func (ms *MemoryStore) GetUser(id uuid.UUID) (User, error) {
	user, exists := ms.users[id.String()]
	if !exists {
		return User{}, NotFound{Resource: "User", Id: id}
	}

	return user, nil
}

func (ms *MemoryStore) GetAllUsers() ([]User, error) {
	res := []User{}
	for _, user := range ms.users {
		res = append(res, user)
	}

	return res, nil
}

func (ms *MemoryStore) CreateSecret(newSecret NewSecret) (Secret, error) {
	secret := newSecret.ToSecret()
	ms.secrets[secret.Id.String()] = secret

	return secret, nil
}

func (ms *MemoryStore) GetSecret(id uuid.UUID) (Secret, error) {
	secret, exists := ms.secrets[id.String()]
	if !exists {
		return Secret{}, NotFound{Resource: "Secret", Id: id}
	}

	return secret, nil
}

func (ms *MemoryStore) GetAllSecrets() ([]Secret, error) {
	res := []Secret{}
	for _, secret := range ms.secrets {
		res = append(res, secret)
	}

	return res, nil
}

func (ms *MemoryStore) UpdateSecret(id uuid.UUID, guesser uuid.UUID) (Secret, error) {
	secret, err := ms.GetSecret(id)
	if err != nil {
		return Secret{}, err
	}

	secret.Guessed = true
	secret.GuesserId = guesser

	ms.secrets[secret.Id.String()] = secret

	return secret, nil
}
